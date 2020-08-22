function redirect(id) {
    let s = "http://" + location.hostname + ":" + location.port + "/todo/" + id;
    console.log(s)
    return s;
}


function update(id) {
    let title = document.getElementById("title")
    let titleInput = document.getElementById("title-input")
    let description = document.getElementById("description")
    let descriptionInput = document.getElementById("description-input")
    let status = document.getElementById("status")
    let statusInput = document.getElementById("status-input")
    let actionButton = document.getElementById("action-button")

    if (actionButton.innerText === "edit"){
        title.style.cssText = "display: none;"
        description.style.cssText = "display: none;"
        status.style.cssText = "display: none;"

        actionButton.innerText = "save"

        titleInput.style.cssText = "display: block;"
        descriptionInput.style.cssText = "display: block;"
        statusInput.style.cssText = "display: block;"
        return
    }else {
        title.style.cssText = "display: block;"
        description.style.cssText = "display: block;"
        status.style.cssText = "display: block;"

        titleInput.style.cssText = "display: none;"
        descriptionInput.style.cssText = "display: none;"
        statusInput.style.cssText = "display: none;"

        let request = new XMLHttpRequest();
        request.open("PUT", "http://localhost:8080/todo/" + id, false)
        let body = JSON.stringify({
            "title": titleInput.value,
            "description": descriptionInput.value,
            "status": statusInput.value === "true"
        });

        request.send(body)

        window.location.href = redirect(id)
    }
}