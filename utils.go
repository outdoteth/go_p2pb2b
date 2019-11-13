package p2pb2b

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/// A general utility for making unauthenticated api requests
func (clt *Client) APIRequest(method, endpoint string) ([]byte, error) {
	ctx := clt.Ctx
	var req *http.Request
	var err error
	if method == http.MethodGet {
		req, err = http.NewRequest(method, clt.URL+endpoint, nil)
	}

	if method == http.MethodPost {
		//run through all of the signing procedures
	}

	req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(body))
	return body, nil
}

/// A general utility for making authenticated API requests
/// Using HMACSha512 signing
func (clt *Client) AuthAPIRequest(postDataBuffer interface{}, method, endpoint string) ([]byte, error) {
	ctx := clt.Ctx
	var req *http.Request
	var err error

	// Gets the json body ready for signing
	byteJson, err := json.Marshal(postBody)
	if err != nil {
		return nil, err
	}
	payload := base64.StdEncoding.EncodeToString(byteJson)

	/// Signs a sha512 hash of the json base64 string and
	/// returns a hex string of the signature
	h := hmac.New(sha512.New, []byte(clt.APIKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	//run through all of the signing procedures
	/*
			body { request: <url><string>
				nonce: Date.now()
				...other_data
			}

		1. get the postData (should already be formatted to a string/buffer) then make it into a body

		signature = toString(body)
			.toBuffer()
			.toBase64()
			.Sha512HMACSign()
			.toHex()

		request {
			url: completeURL,
			  headers: {
			    'Content-Type': 'application/json',
			    'X-TXC-APIKEY': apiKey,
			    'X-TXC-PAYLOAD': toString(body).Buffer.toBase64(),
			    'X-TXC-SIGNATURE': signature
			  },
			  body: toString(body),
		}
	*/

	req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(body))
	return body, nil
}
