import { render, screen } from "@testing-library/react";
import {  User } from "../../interface";
import axios from "axios";
import { UserPage } from "../UserPage";

describe("Userコンポーネント", () => {
  it("user page", async() => {
    const res = await axios.get(`http://127.0.0.1:4000/api/users/1`, { withCredentials: true, adapter: require('axios/lib/adapters/http') })
    const user = res.data[0] as User
    expect(res.status).toBe(200)
    expect(user.id).toBe(1)
    render(<UserPage user={user} me={false}  />);

    expect(screen.getByTestId('bio')).toBe(user.bio)
  });
});