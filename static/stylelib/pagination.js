function redirect(offset, limit) {
    return "http://" + location.hostname + ":" + location.port + "/todo?offset=" + offset + "&limit=" + limit;
}

function changePage(elem,selectedLimit) {
    let urlParams = new URLSearchParams(window.location.search);
    let offset = parseInt(urlParams.get('offset'));
    let limit = parseInt(urlParams.get('limit'));

    if (selectedLimit != null){
        limit = selectedLimit
    }

    if (elem.innerHTML === "Back") {
        if (offset !== 0) {
            offset -= 1
        }
    }

    if (elem.innerHTML === "Next") {
        offset += 1
    }

    console.log("offset -> " + offset)
    console.log("limit -> " + limit)

    window.location.href = redirect(offset, limit)
}