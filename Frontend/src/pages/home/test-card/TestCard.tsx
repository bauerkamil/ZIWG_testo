import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import { ITestsResponse } from "@/shared/interfaces/ITestsResponse";
import { formatDate } from "@/shared/utils/helpers";
import { useNavigate } from "react-router-dom";

const TestCard = (props: { test: ITestsResponse }) => {
  const { test } = props;
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/test/${test.id}`);
  }

  return (
    <div onClick={handleClick}>
      <Card className="cursor-pointer hover:border-solid hover:border-current group">
        <CardHeader>
          <div className="font-semibold leading-none tracking-tight group-hover:text-2xl">
            {test.name}
          </div>
        </CardHeader>
        <CardContent>
          <div className="flex flex-column space-x-1">
            <div className="text-primary font-bold">{test.course.name}</div>
            <div className="text-primary font-bold">{test.course.school_year}</div></div>
          <div className="flex flex-row">
            <p className="text-muted-foreground italic">Prowadzący:&nbsp;</p>
            <p>{test.course.teacher.name} {test.course.teacher.surname}</p>
          </div>
        </CardContent>
        <CardFooter>
          <div className="text-sm text-muted-foreground">
            <div>
              Stworzone dnia {formatDate(test.createdAt)}
            </div>
            <div>
              Twórca: {test.createdBy}
            </div>
            <div>
              Ostatnia modyfikacja: {test.changedAt ? formatDate(test.changedAt) : formatDate(test.createdAt)}
            </div>
          </div>
        </CardFooter>
      </Card>
    </div>
  );
};

export default TestCard;
