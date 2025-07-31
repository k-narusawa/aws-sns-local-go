const API_BASE_URL = import.meta.env.VITE_API_URL ?? "";

export interface Topic {
  topic_arn: string;
  attributes: string;
  tags: string;
}

export interface TopicsResponse {
  items: Topic[];
}

export interface Message {
  message_id: string;
  topic_arn: string;
  target_arn: string;
  phone_number: string;
  message: string;
  subject: string;
  message_structure: string;
  message_attributes: string;
  message_deduplication_id: string;
  message_group_id: string;
  published_at: string;
}

export interface MessagesResponse {
  page: number;
  limit: number;
  size: number;
  totalPage: number;
  totalSize: number;
  items: Message[];
}

export const createTopic = async (name: string): Promise<Topic> => {
  const response = await fetch(`${API_BASE_URL}/topics`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ name }),
  });
  return response.json();
};

export const getTopics = async (): Promise<Topic[]> => {
  const response = await fetch(`${API_BASE_URL}/topics`);
  if (!response.ok) {
    throw new Error(`Failed to fetch topics: ${response.statusText}`);
  }
  const data: TopicsResponse = await response.json();
  return data.items || [];
};

export const getMessages = async (topicArn?: string): Promise<Message[]> => {
  const url = topicArn
    ? `${API_BASE_URL}/messages?topic_arn=${encodeURIComponent(topicArn)}`
    : `${API_BASE_URL}/messages`;
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`Failed to fetch messages: ${response.statusText}`);
  }
  const data: MessagesResponse = await response.json();
  return data.items || [];
};
