import { IAnswear } from "@/shared/interfaces";

export interface ISolveAnswearProps {
  answear: IAnswear;
  revealed?: boolean;
  selected?: boolean;
  onSelected: (solvedAnswear: IAnswear) => void;
}
