const APP_HOST = "http://localhost:4080/"
const user = "admin"
const pass = "Complexpass#123"
const base64Credentials = btoa(user + ':' + pass);

export const getData = async function (url = "", params = {}) {
    url = APP_HOST + url;
    if (params !== {}) {
        url += "?" + new URLSearchParams(params).toString();
    }
    const response = await fetch(url, {
        method: "GET",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Basic " + base64Credentials
        },
        redirect: "follow",
        referrerPolicy: "no-referrer",
    });
    return response.json();
};

export const postData = async function (url = "", data = {}) {
    url = APP_HOST + url;
    const response = await fetch(url, {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Basic " + base64Credentials
        },
        redirect: "follow",
        referrerPolicy: "no-referrer",
        body: JSON.stringify(data),
    });
    return response.json();
};