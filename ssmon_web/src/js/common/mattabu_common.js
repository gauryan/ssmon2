import CryptoJS from "crypto-js";


let SESSION_CRYPTO_KEY = "SESSION_CRYPTO_KEY";


/**
 * SESSION_CRYPTO_KEY 를 설정한다.
 * 보통 로그인 성공한 후에 수행한다.
 */
export function set_session_crypto_key(key) {
    SESSION_CRYPTO_KEY = key;
}


/**
 * 세션스토리지의 값을 설정한다. (암호화)
 */
export function g_set_crypto_session(key, value) {
    const enc_value = CryptoJS.AES.encrypt(value, SESSION_CRYPTO_KEY).toString();
    sessionStorage.setItem(key, enc_value);
}


/**
 * 세션스토리지의 값을 가져온다. (복호화)
 */
export function g_get_crypto_session(key) {
    const enc_value = sessionStorage.getItem(key);
    const decrypted_bytes = CryptoJS.AES.decrypt(enc_value, SESSION_CRYPTO_KEY);
    const decrypted = decrypted_bytes.toString(CryptoJS.enc.Utf8);
    return decrypted;
}

