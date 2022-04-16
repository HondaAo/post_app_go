export interface User {
    id: number,
    created_at: string,
    updated_at: string,
    username: string,
    email: string,
    bio: string,
    posts: PostProps[],
    votes: Vote[]
}

export interface PostProps {
    id: number,
    created_at: string,
    updated_at: string,
    points: number,
    text: string,
    title: string,
    user_id: number,
    replies: null,
    votes: null
}
export interface Vote {
    id: number,
    created_at: string,
    updated_at: string, 
    user_id: number,
    post_id: number,
    value: number
}