import { Button, LinkButton } from "@/components/ui";
import QuestionSummary from "../question-summary/QuestionSummary";
import { ITestSummaryProps } from "./ITestSummaryProps";
import styles from "./TestSummary.module.scss";
import { ArrowDown } from "lucide-react";
import { useEffect, useState } from "react";

const TestSummary = (props: ITestSummaryProps) => {
  const { questions, solvedQuestions, onRefresh } = props;
  const [isScrollVisible, setScrollIsVisible] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
        setScrollIsVisible(false);
      } else {
        setScrollIsVisible(true);
      }
    };

    window.addEventListener('scroll', handleScroll);
    
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, []);

  const handleScrollToBottom = () => {
    window.scrollTo({
      top: document.body.scrollHeight,
      behavior: 'smooth',
    });
  };

  return (
    <div className="grid gap-4">
      {solvedQuestions.map((solvedQuestion, i) => (
        <QuestionSummary
          key={i}
          questionSolved={solvedQuestion}
          question={questions.find((x) => x.id === solvedQuestion.id)}
        />
      ))}
      <div className="flex flex-col lg:flex-row gap-2">
        <div className="grow"></div>
        <LinkButton href="/home" variant="secondary">
          Wróć do widoku głównego
        </LinkButton>
        <Button onClick={onRefresh}>Rozwiąż test jeszcze raz</Button>
      </div>
      {isScrollVisible && <Button 
        className={styles.scrollButton} 
        onClick={handleScrollToBottom}
        >
          <ArrowDown />
        </Button>}
    </div>
  );
};

export default TestSummary;
