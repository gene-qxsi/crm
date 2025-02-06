document.getElementById("register-form").addEventListener("submit", async (event) => {
    event.preventDefault(); // Предотвращаем перезагрузку страницы

    const formData = new FormData(event.target);
    const data = Object.fromEntries(formData.entries()); // Преобразуем в JSON-объект

    try {
        const response = await fetch("/users", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data),
        });

        const text = await response.text();
        alert(text);
    } catch (error) {
        console.error("Ошибка:", error);
        alert("Ошибка отправки данных.");
    }
});