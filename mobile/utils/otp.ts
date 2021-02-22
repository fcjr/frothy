import jsSHA from 'jssha'
import base32 from 'hi-base32'

export type TOTP = {
	code: string,
	expiresAt: Date
}

export const genTOTP = (secret: string): TOTP => {
	const hashFunc = 'SHA-1'   // TODO support mutliple hash functions
	const intervalSeconds = 30 // TODO support multiple intervals
	const codeLen = 6          // TODO support multiple code lengths

	const startTime = Math.floor(Date.now() / 1000 / intervalSeconds)
	const counter = startTime.toString(16).toUpperCase().padStart(16, '0')

	const hmac = new jsSHA(
		hashFunc,
		'HEX',
		{ hmacKey: { value: base32.decode(secret, true), format: 'BYTES' } },
	)
	hmac.update(counter)
	const sum = hmac.getHash('UINT8ARRAY')

	const offset = sum[19] & 0xf
	const binCode = ((sum[offset] & 0x7f) << 24 |
		((sum[offset+1] & 0xff) << 16) |
		((sum[offset+2] & 0xff) << 8) |
		(sum[offset+3] & 0xff))

	const code = binCode % 10 ** codeLen
	const paddedCode = code.toString().padStart(codeLen, '0')

	return {
		code: paddedCode,
		expiresAt: new Date((startTime+1) * 1000 * intervalSeconds)
	}
}
