function changeStatus(elem, id) {
    console.log(elem.checked)
    let request = new XMLHttpRequest();
    request.open("PATCH", "http://localhost:8080/todo/" + id + "?status=" + elem.checked, false)
    request.send(null)

    location.reload()
}