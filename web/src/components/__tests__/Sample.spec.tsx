import { Sample } from "../Sample";
import { render } from "@testing-library/react";

describe("Sampleコンポーネント", () => {
  it("post test", () => {
    const { getByText } = render(<Sample />);
    expect(getByText("Nextjs+Jestのサンプルサプリ")).toBeTruthy();
    expect(getByText("設定がすごく楽になりました。")).toBeTruthy();
  });
});