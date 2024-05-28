import axios from 'axios';
export let backUrl = process.env.back_end_url ?? 'https://tg-game-production-8e6f.up.railway.app'

export function GetWallet(userId, walletId) {
    const data = {
        user_id: userId,
        wallet_id: walletId
    }
    let query = new URLSearchParams(data);
    let fullUrl = backUrl + "/wallet" + query.toString()
    const response = axios.get(fullUrl)
    return response
}

export async function CreateWallet(userId) {
    const data = {
        user_id: userId,
    }
    let query = new URLSearchParams(data);
    let fullUrl = backUrl + "/wallet" + query.toString()
    const s = await axios.post(fullUrl)
    return s
}

export function GetUser(userId) {
    const data = {
        user_id: userId,
    }
    let query = new URLSearchParams(data);
    let fullUrl = backUrl + "/user" + query.toString()
    const response = axios.get(fullUrl)
    return response
}
