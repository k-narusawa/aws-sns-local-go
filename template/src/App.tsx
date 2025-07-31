import { useState, useEffect, useCallback } from "react";
import "./App.css";
import {
  Topic,
  Message,
  createTopic,
  getTopics,
  publishMessage,
  getMessages,
} from "./api/sns";

function App() {
  const [topics, setTopics] = useState<Topic[]>([]);
  const [messages, setMessages] = useState<Message[]>([]);
  const [newTopic, setNewTopic] = useState("");
  const [selectedTopic, setSelectedTopic] = useState<Topic | null>(null);
  const [newMessage, setNewMessage] = useState("");
  const [error, setError] = useState<string | null>(null);

  const fetchTopics = useCallback(async () => {
    try {
      console.log("Fetching topics...");
      const fetchedTopics = await getTopics();
      console.log("Fetched topics:", fetchedTopics);
      setTopics(fetchedTopics || []);
    } catch (err) {
      setError("トピックの取得に失敗しました");
      console.error("Error fetching topics:", err);
    }
  }, []);

  const fetchMessages = useCallback(async () => {
    try {
      console.log("Fetching messages...");
      const fetchedMessages = await getMessages(selectedTopic?.arn);
      console.log("Fetched messages:", fetchedMessages);
      setMessages(fetchedMessages || []);
    } catch (err) {
      setError("メッセージの取得に失敗しました");
      console.error("Error fetching messages:", err);
    }
  }, [selectedTopic?.arn]);

  useEffect(() => {
    fetchTopics();
    const topicsInterval = setInterval(fetchTopics, 5000);
    return () => clearInterval(topicsInterval);
  }, [fetchTopics]);

  useEffect(() => {
    fetchMessages();
    const messagesInterval = setInterval(fetchMessages, 3000);
    return () => clearInterval(messagesInterval);
  }, [fetchMessages, selectedTopic]);

  const handleCreateTopic = async () => {
    if (newTopic.trim()) {
      try {
        await createTopic(newTopic.trim());
        setNewTopic("");
        fetchTopics();
      } catch (err) {
        setError("トピックの作成に失敗しました");
        console.error(err);
      }
    }
  };

  const handleSendMessage = async () => {
    if (selectedTopic && newMessage.trim()) {
      try {
        await publishMessage(selectedTopic.arn, newMessage.trim());
        setNewMessage("");
        fetchMessages();
      } catch (err) {
        setError("メッセージの送信に失敗しました");
        console.error(err);
      }
    }
  };

  return (
    <div className="container">
      <header className="header">
        <h1>AWS SNS Local</h1>
        <p className="version">v0.3.0</p>
      </header>

      <main className="main">
        {error && <div className="error-message">{error}</div>}

        <section className="topics-section">
          <h2>Topics</h2>
          <div className="topic-form">
            <input
              type="text"
              value={newTopic}
              onChange={(e) => setNewTopic(e.target.value)}
              placeholder="トピック名を入力"
            />
            <button onClick={handleCreateTopic}>トピックを作成</button>
          </div>
          <ul className="topics-list">
            {topics.map((topic) => (
              <li
                key={topic.arn}
                className={`topic-item ${
                  selectedTopic?.arn === topic.arn ? "selected" : ""
                }`}
                onClick={() => setSelectedTopic(topic)}
              >
                <span className="topic-name">{topic.name}</span>
                <span className="topic-arn">{topic.arn}</span>
              </li>
            ))}
          </ul>
        </section>

        <section className="messages-section">
          <h2>Messages</h2>
          <div className="message-form">
            <input
              type="text"
              value={newMessage}
              onChange={(e) => setNewMessage(e.target.value)}
              placeholder="メッセージを入力"
              disabled={!selectedTopic}
            />
            <button onClick={handleSendMessage} disabled={!selectedTopic}>
              メッセージを送信
            </button>
          </div>
          <ul className="messages-list">
            {messages.map((msg) => (
              <li key={msg.message_id} className="message-item">
                <div className="message-header">
                  <span className="message-id">ID: {msg.message_id}</span>
                  <span className="message-date">
                    {new Date(msg.published_at).toLocaleString()}
                  </span>
                </div>
                <div className="message-details">
                  {msg.topic_arn && (
                    <span className="message-topic-arn">
                      Topic: {msg.topic_arn}
                    </span>
                  )}
                  {msg.phone_number && (
                    <span className="message-phone">
                      Phone: {msg.phone_number}
                    </span>
                  )}
                  {msg.subject && (
                    <span className="message-subject">
                      Subject: {msg.subject}
                    </span>
                  )}
                </div>
                <span className="message-content">{msg.message}</span>
              </li>
            ))}
          </ul>
        </section>
      </main>
    </div>
  );
}

export default App;
