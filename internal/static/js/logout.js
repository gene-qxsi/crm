document.getElementById("logoutForm").addEventListener("submit", function(event) {
    event.preventDefault()
    
    const op = "internal.pages.js.registration.addEventListener"
    fetch("/api/logout", {
        method: "POST"
    })
    .then(response => response.json())
    .then(data => {
        console.log(data.message, data.status)
        if (data.status === "success") {
            window.location.href = data.redirect
        } else {
            console.log("❌ SERVER-ERROR:", data.error, "PATH:", op)
        }
    })
    .catch(error => console.error("❌ JS-ERROR:", error, "PATH:", op))
})