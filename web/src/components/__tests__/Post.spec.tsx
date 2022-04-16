import { PostComponent }  from "../Post";
import { fireEvent, render, screen } from "@testing-library/react";
import { PostProps } from "../../interface";
import axios from "axios";

interface PostComponentProps {
    post: PostProps
}
describe("Postコンポーネント", () => {
  it("should first", async() => {
    const res = await axios.get(`http://127.0.0.1:4000/api/posts`, { withCredentials: true, adapter: require('axios/lib/adapters/http') })
    const post = res.data[0] as PostProps
    const oldPoints = post.points
    expect(res.status).toBe(200)
    expect(post.user_id).toBe(1)
    render(<PostComponent post={post}  />);

    const upvote = screen.getByTestId('upvote')
    upvote.click()

    expect(post.points).toBe(oldPoints + 1)

  });
});