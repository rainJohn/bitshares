// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: callorderupdateoperation.go

package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *CallOrderUpdateOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *CallOrderUpdateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	/* Struct fall back. type=objects.AssetAmount kind=struct */
	buf.WriteString(`{"delta_collateral":`)
	err = buf.Encode(&j.DeltaCollateral)
	if err != nil {
		return err
	}
	/* Struct fall back. type=objects.AssetAmount kind=struct */
	buf.WriteString(`,"delta_debt":`)
	err = buf.Encode(&j.DeltaDebt)
	if err != nil {
		return err
	}
	buf.WriteString(`,"funding_account":`)

	{

		obj, err = j.FundingAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=objects.AssetAmount kind=struct */
	buf.WriteString(`,"fee":`)
	err = buf.Encode(&j.Fee)
	if err != nil {
		return err
	}
	buf.WriteString(`,"extensions":`)
	if j.Extensions != nil {
		buf.WriteString(`[`)
		for i, v := range j.Extensions {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Interface types must use runtime reflection. type=interface {} kind=interface */
			err = buf.Encode(v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtCallOrderUpdateOperationbase = iota
	ffjtCallOrderUpdateOperationnosuchkey

	ffjtCallOrderUpdateOperationDeltaCollateral

	ffjtCallOrderUpdateOperationDeltaDebt

	ffjtCallOrderUpdateOperationFundingAccount

	ffjtCallOrderUpdateOperationFee

	ffjtCallOrderUpdateOperationExtensions
)

var ffjKeyCallOrderUpdateOperationDeltaCollateral = []byte("delta_collateral")

var ffjKeyCallOrderUpdateOperationDeltaDebt = []byte("delta_debt")

var ffjKeyCallOrderUpdateOperationFundingAccount = []byte("funding_account")

var ffjKeyCallOrderUpdateOperationFee = []byte("fee")

var ffjKeyCallOrderUpdateOperationExtensions = []byte("extensions")

// UnmarshalJSON umarshall json - template of ffjson
func (j *CallOrderUpdateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *CallOrderUpdateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtCallOrderUpdateOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtCallOrderUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffjKeyCallOrderUpdateOperationDeltaCollateral, kn) {
						currentKey = ffjtCallOrderUpdateOperationDeltaCollateral
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyCallOrderUpdateOperationDeltaDebt, kn) {
						currentKey = ffjtCallOrderUpdateOperationDeltaDebt
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyCallOrderUpdateOperationExtensions, kn) {
						currentKey = ffjtCallOrderUpdateOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyCallOrderUpdateOperationFundingAccount, kn) {
						currentKey = ffjtCallOrderUpdateOperationFundingAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyCallOrderUpdateOperationFee, kn) {
						currentKey = ffjtCallOrderUpdateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyCallOrderUpdateOperationExtensions, kn) {
					currentKey = ffjtCallOrderUpdateOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCallOrderUpdateOperationFee, kn) {
					currentKey = ffjtCallOrderUpdateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCallOrderUpdateOperationFundingAccount, kn) {
					currentKey = ffjtCallOrderUpdateOperationFundingAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCallOrderUpdateOperationDeltaDebt, kn) {
					currentKey = ffjtCallOrderUpdateOperationDeltaDebt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCallOrderUpdateOperationDeltaCollateral, kn) {
					currentKey = ffjtCallOrderUpdateOperationDeltaCollateral
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtCallOrderUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtCallOrderUpdateOperationDeltaCollateral:
					goto handle_DeltaCollateral

				case ffjtCallOrderUpdateOperationDeltaDebt:
					goto handle_DeltaDebt

				case ffjtCallOrderUpdateOperationFundingAccount:
					goto handle_FundingAccount

				case ffjtCallOrderUpdateOperationFee:
					goto handle_Fee

				case ffjtCallOrderUpdateOperationExtensions:
					goto handle_Extensions

				case ffjtCallOrderUpdateOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_DeltaCollateral:

	/* handler: j.DeltaCollateral type=objects.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=objects.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.DeltaCollateral)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DeltaDebt:

	/* handler: j.DeltaDebt type=objects.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=objects.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.DeltaDebt)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FundingAccount:

	/* handler: j.FundingAccount type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.FundingAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=objects.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=objects.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=objects.Extensions kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Extensions", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Extensions = nil
		} else {

			j.Extensions = []interface{}{}

			wantVal := true

			for {

				var tmpJExtensions interface{}

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJExtensions type=interface {} kind=interface quoted=false*/

				{
					/* Falling back. type=interface {} kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJExtensions)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.Extensions = append(j.Extensions, tmpJExtensions)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
