import React, { useState } from 'react';
import './Chatbot.css';
import logo from '../../images/chatbot-icon.png';

function Chatbot() {
  const [isOpen, setIsOpen] = useState(false);
  const [message, setMessage] = useState('');
  const [chatLog, setChatLog] = useState([]);

  const toggleChatbot = () => {
    setIsOpen(!isOpen);
  };

  const handleInputChange = (event) => {
    setMessage(event.target.value);
  };

  const sendMessage = () => {
    if (message.trim() === '') return;

    // 添加新消息到聊天记录
    setChatLog([...chatLog, message]);

    // 清空输入框
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
