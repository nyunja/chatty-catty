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
            messageItem.classList.add("my-message");
        };
        messageItem.innerHTML = `
            <div class="message-content">
                <h4>${sender}</h4>
                <p>${message}</p>
            </div>
        `;
        chatArea.appendChild(messageItem);
        chatArea.scrollTop = chatArea.scrollHeight; // scroll to bottom
    };

    // Event handler for form submission
    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const messageText = input.ariaValueMax.trim();
        if (!messageText) return; // skip empty messages

        // Append users message to the chat area
        appendMessage("John", messageText);

        // Send message to the server
        try {
            const response = await fetch('/chat', {
                method: "POST",
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({message: messageText})
            });

            if (response.ok) {
                const data = await response.json();
                appendMessage("Jill", data.response)
            } else {
                console.error('Server error:', data.response)
                appendMessage('Jill', 'Oop! I didnt get you clearly. Please rephrase')
            }

        } catch (error) {
            console.error('Fetch error:', error)
            appendMessage('Jill', 'Error: Connecion issue')
        }
        input.value = '';
    });

    // Click event to sedn message when send icon is clicked
    sendIcon.addEventListener('click', () => form.dispatchEvent(new Event('submit')));
});