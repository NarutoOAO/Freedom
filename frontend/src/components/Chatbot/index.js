import React, { useState } from 'react';
import './Chatbot.css';
import logo from '../../images/chatbot-icon.png';
// define the chatbot
function Chatbot() {
  const [isOpen, setIsOpen] = useState(false);
  const [message, setMessage] = useState('');
  const [chatLog, setChatLog] = useState([]);
// check the isOpen state and toggle it
  const toggleChatbot = () => {
    setIsOpen(!isOpen);
  };
// handle the input change
  const handleInputChange = (event) => {
    setMessage(event.target.value);
  };
//  handle the send message
  const sendMessage = () => {
    if (message.trim() === '') return;

//  set the chat log
    setChatLog([...chatLog, message]);

// clear the input
    setMessage('');
  };

  return (
    <div className={`chatbot-container ${isOpen ? 'open' : ''}`}>
      {isOpen && (
        <div className="chatbot-window">
          <div className="chat-log">
            {chatLog.map((message, index) => (
              <div key={index} className="message">
                {message}
              </div>
            ))}
          </div>
          <div className="input-container">
            <input
              type="text"
              value={message}
              onChange={handleInputChange}
              placeholder="Type your message..."
            />
            <button className="send-button" onClick={sendMessage}>
              Send
            </button>
          </div>
        </div>
      )}
      <button className="chatbot-icon" onClick={toggleChatbot}>
        <img src={logo} alt="Chatbot Icon" />
      </button>
    </div>
  );
}

export default Chatbot;
