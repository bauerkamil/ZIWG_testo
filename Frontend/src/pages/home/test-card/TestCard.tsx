import { useState } from 'react';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import { ITest } from "@/shared/interfaces/ITest";
import { formatDate } from "@/shared/utils/helpers";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";

const TestCard = (props: { test: ITest }) => {
  const { test } = props;
  const navigate = useNavigate();
  const [showOptions, setShowOptions] = useState(false);

  const handleClick = () => {
    setShowOptions(true);
  };

  const handleCloseOptions = () => {
    setShowOptions(false);
  };

  const handleSolveClick = () => {
    navigate(`/solve/${test.id}`);
    handleCloseOptions();
  };

  const handleEditClick = () => {
    navigate(`/edit/${test.id}`)
    handleCloseOptions();
  };

  return (
    <div>
      <Card
        className="cursor-pointer hover:border-solid hover:border-current group"
        onClick={handleClick}
      >
        <CardHeader>
          <div className="font-semibold leading-none tracking-tight group-hover:text-2xl">
            {test.name}
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
      {showOptions && (
        <div className="fixed inset-0 z-50 flex items-center justify-center">
          <div className="absolute inset-0 bg-gray-800 opacity-50" onClick={handleCloseOptions}></div>
          <Card className="relative bg-white p-8 rounded shadow-lg flex flex-col items-center justify-center">
            <p className="text-lg mb-4">Co chcesz zrobić z testem "{test.name}"?</p>
            <Button className="mb-4" onClick={handleSolveClick}>Rozwiąż test</Button>
            <Button onClick={handleEditClick}>Edytuj test</Button>
          </Card>
        </div>
      )}
    </div>
  );
};

export default TestCard;
