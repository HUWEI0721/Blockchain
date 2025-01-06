import request from '@/utils/request'

export function uplink(data) {
  return request({
    url: '/uplink',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getSeafoodInfo
export function getSeafoodInfo(data) {
  return request({
    url: '/getSeafoodInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getSeafoodList
export function getSeafoodList(data) {
  return request({
    url: '/getSeafoodList',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getAllSeafoodInfo
export function getAllSeafoodInfo(data) {
  return request({
    url: '/getAllSeafoodInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getSeafoodHistory
export function getSeafoodHistory(data) {
  return request({
    url: '/getSeafoodHistory',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

