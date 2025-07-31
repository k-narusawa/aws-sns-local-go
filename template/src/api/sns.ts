export interface Topic {
  name: string;
  arn: string;
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
  const response = await fetch(`/topics`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ name }),
  });
  if (!response.ok) {
    throw new Error(`Failed to create topic: ${response.statusText}`);
  }
  return response.json();
};

export const getTopics = async (): Promise<Topic[]> => {
  const response = await fetch(`/topics`);
  if (!response.ok) {
    throw new Error(`Failed to fetch topics: ${response.statusText}`);
  }
  const data = await response.json();
  return Array.isArray(data) ? data : [];
};

export const publishMessage = async (
  topicArn: string,
  message: string
): Promise<void> => {
  const response = await fetch(`/messages`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      topic_arn: topicArn,
      message: message,
    }),
  });
  if (!response.ok) {
    throw new Error(`Failed to publish message: ${response.statusText}`);
  }
};

export const getMessages = async (topicArn?: string): Promise<Message[]> => {
  const url = topicArn
    ? `/messages?topic_arn=${encodeURIComponent(topicArn)}`
    : `/messages`;
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`Failed to fetch messages: ${response.statusText}`);
  }
  const data: MessagesResponse = await response.json();
  return data.items || [];
};
