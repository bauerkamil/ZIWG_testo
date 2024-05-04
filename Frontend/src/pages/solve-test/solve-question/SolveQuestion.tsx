import React from "react";
import { shuffle } from "@/shared/utils/helpers";
import { ISolveQuestionProps } from "./ISolveQuestionProps";
import SolveAnswear from "./solve-answear/SolveAnswear";
import { IAnswear } from "@/shared/interfaces";
import { Button } from "@/components/ui";
import { ArrowRight } from "lucide-react";

const SolveQuestion = (props: ISolveQuestionProps) => {
  const { question, onNext, onSkip } = props;
  const [answears, setAnswears] = React.useState<IAnswear[]>([]);
  const [selectedAnswears, setSelectedAnswears] = React.useState<IAnswear[]>(
    []
  );

  const [showAll, setShowAll] = React.useState(false);
  const [nextDisabled, setNextDisabled] = React.useState(false);

  const handleAnswearSelected = (solvedAnswear: IAnswear) => {
    selectedAnswears.push(solvedAnswear);
  };

  const finish = (): Promise<void> => {
    return new Promise((resolve) => {
      setShowAll(true);
      setNextDisabled(true);
      setTimeout(() => {
        setShowAll(false);
        setSelectedAnswears([]);
        setNextDisabled(false);
        resolve();
      }, 2000);
    });
  };

  const handleNext = () => {
    const selectedAnswearsCopy = [...selectedAnswears];
    finish().then(() => {
      onNext(selectedAnswearsCopy);
    });
  };
  const handleSkip = () => {
    finish().then(() => {
      onSkip();
    });
  };

  React.useEffect(() => {
    setAnswears(shuffle(question.answears));
  }, [question]);

  return (
    <div className="grid gap-4">
      <div className="text-4xl text-center">{question.body}</div>
      <div>
        {question.imgFile && (
          <img src={question.imgFile} alt="question" className="w-1/2" />
        )}
      </div>
      <div className="flex flex-col gap-1">
        {answears.map((answear) => (
          <SolveAnswear
            key={answear.id}
            answear={answear}
            revealed={showAll}
            selected={selectedAnswears.some(
              (selected) => selected.id === answear.id
            )}
            onSelected={handleAnswearSelected}
          />
        ))}
      </div>
      <div className="flex">
        <div className="grow"></div>
        <Button
          variant={"secondary"}
          disabled={nextDisabled}
          onClick={handleSkip}
        >
          Skip
        </Button>
        <Button disabled={nextDisabled} onClick={handleNext}>
          Kolejne pytanie <ArrowRight className="h-5 w-5" />
        </Button>
      </div>
    </div>
  );
};

export default SolveQuestion;
