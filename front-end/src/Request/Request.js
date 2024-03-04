const APP_HOST = "http://localhost:3000/"

export const getData = async function (url = "", params = {}) {
    url = APP_HOST + url;
    if (params !== null) {
        url += "?" + new URLSearchParams(params).toString();
    }
    const response = await fetch(url, {
        method: "GET",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json",
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
        },
        redirect: "follow",
        referrerPolicy: "no-referrer",
        body: JSON.stringify(data),
    });
    return response.json();
};