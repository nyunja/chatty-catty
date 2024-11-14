document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById("message-input-form");
    const input = document.getElementById("message-input");
    const chatArea = document.getElementById("all-chat");
    const sendIcon = document.getElementById("send-icon");

    // Function to append message to the chat area
    const appendMessage = (sender, message) => {
        const messageItem = document.createElement("div");
        messageItem.classList.add("message-item");
        if (sender.length !== 0) {
            messageItem.classList.add("my-message")
        }
        messageItem.innerHTML = `
            <div class="message-content">
                <h4>${sender}</h4>
                <p>${message}</p>
            </div>
        `;
        chatArea.appendChild(messageItem)
        chatArea.scrollTop = chatArea.scrollHeight; // scroll to bottom
    };

    // Event handler for form submission

    // Click event to sedn message when send icon is clicked
});