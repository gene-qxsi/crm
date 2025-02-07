document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault()

    const op = "internal.pages.js.login.addEventListener"
    const formData = new FormData(event.target)
    const jsonData = {}

    formData.forEach((value, key) => {
        jsonData[key] = isNaN(value) ? value : Number(value)
    })

    fetch("/login", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(jsonData)
    })
    .then(response => response.json())
    .then(data => {
        console.log("Ответ от сервера:", data.message)
        if (data.status === "success") {
            window.location.href = data.redirect
        } else {
            console.log("errrrrrrrrrrrrrorrrrrrrrrr")
        }
    })
    .catch(error => console.error("❌ JS-ERROR:", error, "PATH:", op))
})