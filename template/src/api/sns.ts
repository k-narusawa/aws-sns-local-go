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

export interface CreateTopicRequest {
  Name: string;
  Tags?: string;
  Attributes?: string;
}

export const createTopic = async (
  request: CreateTopicRequest
): Promise<Topic> => {
  const params = new URLSearchParams({
    Action: "CreateTopic",
    ...request,
  });

  const response = await fetch(`${API_BASE_URL}/?${params.toString()}`, {
    method: "POST",
  });

  if (!response.ok) {
    throw new Error(`Failed to create topic: ${response.statusText}`);
  }

  const data = await response.text();
  // XMLレスポンスからTopicArnを抽出
  const match = data.match(/<TopicArn>(.*?)<\/TopicArn>/);
  if (!match) {
    throw new Error("Failed to parse topic ARN from response");
  }

  return {
    topic_arn: match[1],
    attributes: "",
    tags: request.Tags || "",
  };
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
