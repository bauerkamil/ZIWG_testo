import React from "react";
import { ISolveAnswearProps } from "./ISolveAnswearProps";
import { cn } from "@/lib/utils";

const SolveAnswear = (props: ISolveAnswearProps) => {
  const { answear, revealed, selected, onSelected } = props;
  const [answearStyle, setAnswearStyle] = React.useState("");

  const handleClick = () => {
    if (!selected) {
      setAnswearStyle("bg-gray-300");
      onSelected(answear);
    }
  };

  React.useEffect(() => {
    if (!revealed) {
      if (selected) {
        setAnswearStyle("bg-gray-300");
        return;
      }
      setAnswearStyle("");
      return;
    }
    if (answear.valid) {
      setAnswearStyle("bg-green-500");
      return;
    }
    if (selected && !answear.valid) {
      setAnswearStyle("bg-red-500");
      return;
    }
    setAnswearStyle("bg-gray-300");
  }, [answear.valid, revealed, selected]);

  return (
    <div
      className={cn(
        "text-2xl p-2 border border-gray-300 rounded-lg shadow-md transition ease-in-out delay-200",
        answearStyle,
        !revealed && !selected && "cursor-pointer hover:bg-gray-200"
      )}
      onClick={handleClick}
    >
      {answear.body}
    </div>
  );
};

export default SolveAnswear;
