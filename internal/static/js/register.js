document.getElementById("registerForm").addEventListener("submit", function(event) {
    event.preventDefault();

    const op = "internal.pages.js.registration.addEventListener"
    const formData = new FormData(event.target)
    const jsonData = {}

    formData.forEach((value, key) => {
        jsonData[key] = isNaN(value) ? value : Number(value)
    })

    fetch("/api/register", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(jsonData)
    })
    .then(response => response.json())
    .then(data => {
        console.log("Ответ от сервера:", data.message)
        if (data.status === "success") {
            window.location.href = data.redirect
        }
    })
    .catch(error => console.error("❌ JS-ERROR:", error, "PATH:", op))
})