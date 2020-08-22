function remove(id) {
    let request = new XMLHttpRequest();
    request.open("DELETE", "http://localhost:8080/todo/" + id, false)
    let urlParams = new URLSearchParams(window.location.search);
    let body = JSON.stringify({
        "offset": parseInt(urlParams.get('offset')),
        "limit": parseInt(urlParams.get('limit'))
    });
    request.send(body)

    location.reload()
}