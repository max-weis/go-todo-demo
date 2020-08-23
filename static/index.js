const host = "http://" + location.hostname + ":" + location.port;

function changeStatus(elem, id) {
    let request = new XMLHttpRequest();
    request.open("PATCH", host + "/todo/" + id + "?status=" + elem.checked, false)
    request.send(null)

    location.reload()
}

function update(id) {
    let title = document.getElementById("title")
    let titleInput = document.getElementById("title-input")
    let description = document.getElementById("description")
    let descriptionInput = document.getElementById("description-input")
    let status = document.getElementById("status")
    let statusInput = document.getElementById("status-input")
    let actionButton = document.getElementById("action-button")

    if (actionButton.innerText === "edit") {
        title.style.cssText = "display: none;"
        description.style.cssText = "display: none;"
        status.style.cssText = "display: none;"

        actionButton.innerText = "save"

        titleInput.style.cssText = "display: block;"
        descriptionInput.style.cssText = "display: block;"
        statusInput.style.cssText = "display: block;"
        return
    }
    title.style.cssText = "display: block;"
    description.style.cssText = "display: block;"
    status.style.cssText = "display: block;"

    titleInput.style.cssText = "display: none;"
    descriptionInput.style.cssText = "display: none;"
    statusInput.style.cssText = "display: none;"

    let request = new XMLHttpRequest();
    request.open("PUT", host + "/todo/" + id, false)
    let body = JSON.stringify({
        "title": titleInput.value,
        "description": descriptionInput.value,
        "status": statusInput.value === "true"
    });

    request.send(body)

    showDetail(id)
}

function remove(id) {
    let request = new XMLHttpRequest();
    request.open("DELETE", host + "/todo/" + id, false)
    let urlParams = new URLSearchParams(window.location.search);
    let body = JSON.stringify({
        "offset": parseInt(urlParams.get('offset')),
        "limit": parseInt(urlParams.get('limit'))
    });
    request.send(body)

    location.reload()
}

function changePage(elem, selectedLimit) {
    let offset = parseInt(localStorage.getItem("offset"));
    let limit = parseInt(localStorage.getItem("limit"));

    if (selectedLimit != null) {
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

    localStorage.setItem("limit", limit.toString())
    localStorage.setItem("offset", offset.toString())

    showList()
}

function showList() {
    let offset = localStorage.getItem("offset");
    if (offset === null || offset === "") {
        offset = 0
    }

    let limit = localStorage.getItem("limit");
    if (limit === null || limit === "") {
        limit = 5
    }

    window.location.href = host + "/todo?offset=" + offset + "&limit=" + limit
}

function showDetail(id) {
    window.location.href = host + "/todo/" + id
}