import axios from 'axios'
import cache from './cache'
import Lockr from 'lockr'
import store from '@/store'

export function removeAuth() {
    return new Promise((resolve, reject) => {
        cache.rmAxiosCache()
        delete axios.defaults.headers['authToken']
        delete axios.defaults.headers['sessionId']
        resolve(true)
    })
}

export function addAuth(adminToken, sessionId) {
    return new Promise((resolve, reject) => {
        axios.defaults.headers['authToken'] = adminToken
        axios.defaults.headers['sessionId'] = sessionId
        resolve(true)
    })
}

export function getAuth() {
    if (Lockr.get('authToken')) {
        return true
    }
    return false
}