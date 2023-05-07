import request from "../utils/request";

export function wxLogin(data) {
    return request(
        'post',
        '/hl_api/login/wechat',
        data
    )
}

export function userProfile(data) {
    return request(
        'get',
        '/hl_api/user/detail',
        data
    )
}

export function studyrecord(data) {
    return request(
        'get',
        '/hl_api/study/record',
        data
    )
}

export function subjectList(data) {
    return request(
        'get',
        '/hl_api/subject/list',
        data
    )
}

export function study(data) {
    return request(
        'put',
        '/hl_api/subject/study',
        data
    )
}

export function subjectDetail(data) {
    return request(
        'get',
        '/hl_api/subject/detail',
        data
    )
}

export function knowledgeList(data) {
    return request(
        'get',
        '/hl_api/study/knowledge',
        data
    )
}

export function studyDoing(data) {
    return request(
        'post',
        '/hl_api/study/doing',
        data
    )
}