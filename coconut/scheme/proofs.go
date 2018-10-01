// proofs.go - Non-interactive Zero-Knowledge Proofs Implementation
// Copyright (C) 2018  Jedrzej Stuczynski.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package coconut

import (
	"errors"
	"strings"

	"github.com/jstuczyn/CoconutGo/coconut/utils"
	"github.com/jstuczyn/CoconutGo/elgamal"
	"github.com/milagro-crypto/amcl/version3/go/amcl"
	Curve "github.com/milagro-crypto/amcl/version3/go/amcl/BLS381"
)

// todo: Ensure Signer and Verifier are the correct terms for the proofs
// todo: worker pool for concurrency
// todo: make errors private
// todo: deal with too lengthy function signatures

// SignerProof (name to be confirmed) represents all the fields contained within the said proof.
type SignerProof struct {
	c  *Curve.BIG
	rr *Curve.BIG
	rk []*Curve.BIG
	rm []*Curve.BIG
}

// VerifierProof (name to be confirmed) represents all the fields contained within the said proof.
type VerifierProof struct {
	c  *Curve.BIG
	rm []*Curve.BIG
	rt *Curve.BIG
}

// Printable is a wrapper for all objects that have ToString method. In particular Curve.ECP and Curve.ECP2.
type Printable interface {
	ToString() string
}

var (
	// ErrConstructSignerCiphertexts indicates that invalid ciphertexts were provided for construction of
	// proofs for corectness of ciphertexts and cm.
	ErrConstructSignerCiphertexts = errors.New("Invalid ciphertexts provided")

	// ErrConstructSignerAttrs indicates that invalid attributes (either attributes to sign or params generated at setup)
	// were provided for construction of proofs for corectness of ciphertexts and cm.
	ErrConstructSignerAttrs = errors.New("More than specified number of attributes provided")
)

// constructChallenge construct a BIG num challenge by hashing a number of Eliptic Curve points
// It's based on the original Python implementation:
// https://github.com/asonnino/coconut/blob/master/coconut/proofs.py#L9.
func constructChallenge(elems []Printable) *Curve.BIG {
	csa := make([]string, len(elems))
	for i := range elems {
		csa[i] = elems[i].ToString()
	}
	cs := strings.Join(csa, ",")
	c, err := utils.HashStringToBig(amcl.SHA256, cs)
	if err != nil {
		panic(err)
	}
	return c
}

// ConstructSignerProof creates a non-interactive zero-knowledge proof to prove corectness of ciphertexts and cm.
// It's based on the original Python implementation:
// https://github.com/asonnino/coconut/blob/master/coconut/proofs.py#L16
// nolint: interfacer
func ConstructSignerProof(params *Params, gamma *Curve.ECP, encs []*elgamal.Encryption, cm *Curve.ECP, k []*Curve.BIG, r *Curve.BIG, pubM []*Curve.BIG, privM []*Curve.BIG) (*SignerProof, error) {
	attributes := append(privM, pubM...)
	G := params.G
	if len(encs) != len(k) || len(encs) != len(privM) {
		return nil, ErrConstructSignerCiphertexts
	}
	if len(attributes) > len(params.hs) {
		return nil, ErrConstructSignerAttrs
	}

	// witnesses creation
	wr := Curve.Randomnum(G.Ord, G.Rng)
	wk := make([]*Curve.BIG, len(k))
	wm := make([]*Curve.BIG, len(attributes))

	for i := range k {
		wk[i] = Curve.Randomnum(G.Ord, G.Rng)
	}
	for i := range attributes {
		wm[i] = Curve.Randomnum(G.Ord, G.Rng)
	}

	h, err := utils.HashStringToG1(amcl.SHA256, cm.ToString())
	if err != nil {
		return nil, err
	}

	// witnesses commitments
	Aw := make([]*Curve.ECP, len(wk))
	Bw := make([]*Curve.ECP, len(privM))
	var Cw *Curve.ECP

	for i := range wk {
		Aw[i] = Curve.G1mul(G.Gen1, wk[i]) // Aw[i] = (wk[i] * g1)
	}
	for i := range privM {
		Bw[i] = Curve.G1mul(h, wm[i])        // Bw[i] = (wm[i] * h)
		Bw[i].Add(Curve.G1mul(gamma, wk[i])) // Bw[i] = (wm[i] * h) + (wk[i] * gamma)
	}

	Cw = Curve.G1mul(G.Gen1, wr) // Cw = (wr * g1)
	for i := range attributes {
		Cw.Add(Curve.G1mul(params.hs[i], wm[i])) // Cw = (wr * g1) + (wm[0] * hs[0]) + ... + (wm[i] * hs[i])
	}

	tmpSlice := []Printable{G.Gen1, G.Gen2, cm, h, Cw}
	ca := make([]Printable, len(tmpSlice)+len(params.hs)+len(Aw)+len(Bw))
	i := copy(ca, tmpSlice)

	// can't use copy for those due to type difference (Printable vs *Curve.ECP)
	for _, item := range params.hs {
		ca[i] = item
		i++
	}
	for _, item := range Aw {
		ca[i] = item
		i++
	}
	for _, item := range Bw {
		ca[i] = item
		i++
	}

	c := constructChallenge(ca)

	// responses
	rr := wr.Minus(Curve.Modmul(c, r, G.Ord))
	rr.Mod(G.Ord) // rr = (wr - c * r) % o

	rk := make([]*Curve.BIG, len(wk))
	for i := range wk {
		rk[i] = wk[i].Minus(Curve.Modmul(c, k[i], G.Ord))
		rk[i].Mod(G.Ord) // rk[i] = (wk[i] - c * k[i]) % o
	}

	rm := make([]*Curve.BIG, len(wm))
	for i := range wm {
		rm[i] = wm[i].Minus(Curve.Modmul(c, attributes[i], G.Ord))
		rm[i].Mod(G.Ord) // rm[i] = (wm[i] - c * attributes[i]) % o
	}

	return &SignerProof{
			c:  c,
			rr: rr,
			rk: rk,
			rm: rm},
		nil
}

// VerifySignerProof verifies non-interactive zero-knowledge proofs in order to check corectness of ciphertexts and cm.
// It's based on the original Python implementation:
// https://github.com/asonnino/coconut/blob/master/coconut/proofs.py#L41
func VerifySignerProof(params *Params, gamma *Curve.ECP, encs []*elgamal.Encryption, cm *Curve.ECP, proof *SignerProof) bool {
	if len(encs) != len(proof.rk) {
		return false
	}
	G := params.G
	h, err := utils.HashStringToG1(amcl.SHA256, cm.ToString())
	if err != nil {
		panic(err)
	}

	Aw := make([]*Curve.ECP, len(proof.rk))
	Bw := make([]*Curve.ECP, len(encs))
	var Cw *Curve.ECP

	for i := range proof.rk {
		Aw[i] = Curve.G1mul(encs[i].C1(), proof.c)         // Aw[i] = (c * c1[i])
		Aw[i].Add(Curve.G1mul(params.G.Gen1, proof.rk[i])) // Aw[i] = (c * c1[i]) + (rk[i] * g1)
	}

	for i := range encs {
		Bw[i] = Curve.G1mul(encs[i].C2(), proof.c) // Bw[i] = (c * c2[i])
		Bw[i].Add(Curve.G1mul(gamma, proof.rk[i])) // Bw[i] = (c * c2[i]) + (rk[i] * gamma)
		Bw[i].Add(Curve.G1mul(h, proof.rm[i]))     // Bw[i] = (c * c2[i]) + (rk[i] * gamma) + (rm[i] * h)
	}

	Cw = Curve.G1mul(cm, proof.c)                // Cw = (cm * c)
	Cw.Add(Curve.G1mul(params.G.Gen1, proof.rr)) // Cw = (cm * c) + (rr * g1)
	for i := range proof.rm {
		Cw.Add(Curve.G1mul(params.hs[i], proof.rm[i])) // Cw = (cm * c) + (rr * g1) + (rm[0] * hs[0]) + ... + (rm[i] * hs[i])
	}

	tmpSlice := []Printable{G.Gen1, G.Gen2, cm, h, Cw}
	ca := make([]Printable, len(tmpSlice)+len(params.hs)+len(Aw)+len(Bw))
	i := copy(ca, tmpSlice)

	// can't use copy for those due to type difference (Printable vs *Curve.ECP)
	for _, item := range params.hs {
		ca[i] = item
		i++
	}
	for _, item := range Aw {
		ca[i] = item
		i++
	}
	for _, item := range Bw {
		ca[i] = item
		i++
	}

	return Curve.Comp(proof.c, constructChallenge(ca)) == 0
}

// ConstructVerifierProof creates a non-interactive zero-knowledge proof in order to prove corectness of kappa and nu.
// It's based on the original Python implementation:
// https://github.com/asonnino/coconut/blob/master/coconut/proofs.py#L57
func ConstructVerifierProof(params *Params, vk *VerificationKey, sig *Signature, privM []*Curve.BIG, t *Curve.BIG) *VerifierProof {
	G := params.G

	// witnesses creation
	wm := make([]*Curve.BIG, len(privM))
	for i := 0; i < len(privM); i++ {
		wm[i] = Curve.Randomnum(G.Ord, G.Rng)
	}
	wt := Curve.Randomnum(G.Ord, G.Rng)

	// witnesses commitments
	Aw := Curve.G2mul(G.Gen2, wt) // Aw = (wt * g2)
	Aw.Add(vk.alpha)              // Aw = (wt * g2) + alpha
	for i := range privM {
		Aw.Add(Curve.G2mul(vk.beta[i], wm[i])) // Aw = (wt * g2) + alpha + (wm[0] * beta[0]) + ... + (wm[i] * beta[i])
	}
	Bw := Curve.G1mul(sig.sig1, wt) // Bw = wt * h

	tmpSlice := []Printable{G.Gen1, G.Gen2, vk.alpha, Aw, Bw}
	ca := make([]Printable, len(tmpSlice)+len(params.hs)+len(vk.beta))
	i := copy(ca, tmpSlice)

	// can't use copy for those due to type difference (Printable vs *Curve.ECP and *Curve.ECP2)
	for _, item := range params.hs {
		ca[i] = item
		i++
	}
	for _, item := range vk.beta {
		ca[i] = item
		i++
	}

	c := constructChallenge(ca)

	// responses
	rm := make([]*Curve.BIG, len(privM))
	for i := range privM {
		rm[i] = wm[i].Minus(Curve.Modmul(c, privM[i], G.Ord))
		rm[i].Mod(G.Ord)
	}

	rt := wt.Minus(Curve.Modmul(c, t, G.Ord))
	rt.Mod(G.Ord)

	return &VerifierProof{
		c:  c,
		rm: rm,
		rt: rt,
	}
}

// VerifyVerifierProof verifies non-interactive zero-knowledge proofs in order to check corectness of kappa and nu.
// It's based on the original Python implementation:
// https://github.com/asonnino/coconut/blob/master/coconut/proofs.py#L75
func VerifyVerifierProof(params *Params, vk *VerificationKey, sig *Signature, showMats *BlindShowMats) bool {
	G := params.G

	Aw := Curve.G2mul(showMats.kappa, showMats.proof.c) // Aw = (c * kappa)
	Aw.Add(Curve.G2mul(vk.g2, showMats.proof.rt))       // Aw = (c * kappa) + (rt * g2)

	// Aw = (c * kappa) + (rt * g2) + (alpha)
	Aw.Add(vk.alpha)
	// Aw = (c * kappa) + (rt * g2) + (alpha - alpha * c) = (c * kappa) + (rt * g2) + ((1 - c) * alpha)
	Aw.Add(Curve.G2mul(vk.alpha, Curve.Modneg(showMats.proof.c, params.G.Ord)))

	for i := range showMats.proof.rm {
		// Aw = (c * kappa) + (rt * g2) + ((1 - c) * alpha) + (rm[0] * beta[0]) + ... + (rm[i] * beta[i])
		Aw.Add(Curve.G2mul(vk.beta[i], showMats.proof.rm[i]))
	}

	Bw := Curve.G1mul(showMats.nu, showMats.proof.c) // Bw = (c * nu)
	Bw.Add(Curve.G1mul(sig.sig1, showMats.proof.rt)) // Bw = (c * nu) + (rt * h)

	tmpSlice := []Printable{G.Gen1, G.Gen2, vk.alpha, Aw, Bw}
	ca := make([]Printable, len(tmpSlice)+len(params.hs)+len(vk.beta))
	i := copy(ca, tmpSlice)

	// can't use copy for those due to type difference (Printable vs *Curve.ECP and *Curve.ECP2)
	for _, item := range params.hs {
		ca[i] = item
		i++
	}
	for _, item := range vk.beta {
		ca[i] = item
		i++
	}
	return Curve.Comp(showMats.proof.c, constructChallenge(ca)) == 0
}
