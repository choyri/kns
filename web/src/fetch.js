import {Message} from 'element-react'
import moment from 'moment/moment'
import config from './config'

function _fetch(url, options = {}) {
    return fetch(config.BACKEND_URL + url, options)
        .then(resp => {
            if (resp.status < 400) {
                return Promise.resolve(resp)
            }
            return resp.text().then(msg => Promise.reject(msg))
        })
        .catch(err => {
            if (options.returnRawErr) {
                return Promise.reject(err)
            }
            return Promise.reject(Message.error(err))
        })
}

function visitBackend() {
    return _fetch('', {returnRawErr: true})
}

function search(keyword) {
    const url = `/search?q=${encodeURIComponent(keyword)}`

    return _fetch(url)
        .then(resp => resp.json())
}

function exportData(ids) {
    const url = `/export`

    _fetch(url, {
        body: JSON.stringify({'ids': ids}),
        method: 'POST',
    })
        .then(resp => resp.blob())
        .then(blob => {
            const url = window.URL.createObjectURL(blob)

            let a = document.createElement('a')
            a.href = url
            a.download = `订单导出_${moment().format('YMMDD_HHmmss')}.xlsx`

            document.body.appendChild(a)
            a.click()
            a.remove()
        })
}

export {
    visitBackend,
    search,
    exportData,
}
