import request from "../utils/request";

export function wxLogin(data) {
    return request(
        'post',
        '/api/login/wechat',
        data
    )
}

export function userProfile(data) {
    return request(
        'get',
        '/api/user/detail',
        data
    )
}

export function studyrecord(data) {
    return request(
        'get',
        '/api/study/record',
        data
    )
}

export function subjectList(data) {
    return request(
        'get',
        '/api/subject/list',
        data
    )
}

export function study(data) {
    return request(
        'put',
        '/api/subject/study',
        data
    )
}

export function subjectDetail(data) {
    return request(
        'get',
        '/api/subject/detail',
        data
    )
}

export function knowledgeList(data) {
    return request(
        'get',
        '/api/study/knowledge',
        data
    )
}

export function studyDoing(data) {
    return request(
        'post',
        '/api/study/doing',
        data
    )
}