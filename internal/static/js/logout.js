document.getElementById("logoutForm").addEventListener("submit", function(event) {
    event.preventDefault()
    
    fetch("/logout", {
        method: "POST"
    })
    .then(response => response.json())
    .then(data => {
        console.log(data.message, data.status)
        if (data.status === "success") {
            window.location.href = data.redirect
        } else {
            console.log("(t)errrrrrrrrrrrrrrrrrrrrrrrrrorrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
        }
    })
    .catch(error => console.error(error))
})