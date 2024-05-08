import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import { ITest } from "@/shared/interfaces/ITest";
import { formatDate } from "@/shared/utils/helpers";
import { useNavigate } from "react-router-dom";
import TestCardOptions from "./test-card-options/TestCardOptions";

const TestCard = (props: { test: ITest, onDelete: (id: string) => void }) => {
  const { test } = props;
  const { onDelete } = props;
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/solve/${test.id}`);
  };

  return (
    <Card
      className="cursor-pointer hover:border-solid hover:border-current group"
      onClick={handleClick}
    >
      <CardHeader>
        <div className="flex flex-row">
          <div className="font-semibold leading-none tracking-tight content-center group-hover:text-2xl">
            {test.name}
          </div>
          <div className="grow"></div>
          <TestCardOptions id={test.id ?? ""} onDelete={onDelete} />
        </div>
      </CardHeader>
      <CardContent>
        <div className="flex flex-column gap-2">
          <div className="text-primary font-bold">{test.course?.name}</div>
          <div className="text-muted-foreground italic">
            {test.course?.courseType}
          </div>
          {test.schoolYear && (
            <div className="text-muted-foreground italic">
              (Semestr: {test.schoolYear})
            </div>
          )}
        </div>
        <div className="flex flex-row">
          <p className="text-muted-foreground italic">Prowadzący:&nbsp;</p>
          <p>
            {test.course?.teacher.name} {test.course?.teacher.surname}
          </p>
        </div>
      </CardContent>
      <CardFooter>
        <div className="text-sm text-muted-foreground">
          <div>
            Ostatnia modyfikacja:&nbsp;
            {test.changedAt
              ? formatDate(test.changedAt)
              : test.createdAt
                ? formatDate(test.createdAt)
                : "Brak danych"}
          </div>
        </div>
      </CardFooter>
    </Card>
  );
};

export default TestCard;
