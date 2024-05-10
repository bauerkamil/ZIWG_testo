import React from "react";
import Navbar from "@/components/navbar/Navbar";
import { Button, Input, Label, useToast } from "@/components/ui";
import { ICourse, IQuestion, ITest } from "@/shared/interfaces";
import { Plus } from "lucide-react";
import EditQuestion from "./question/EditQuestion";
import { useParams } from "react-router-dom";
import Client from "@/api/Client";
import CourseSelector from "@/components/course-selector/CourseSelector";

const EditTest: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [courses, setCourses] = React.useState<ICourse[]>([]);
  const [test, setTest] = React.useState<ITest>();

  const { toast } = useToast();

  React.useEffect(() => {
    const fetchData = async () => {
      if (id) {
        try {
          const testData = await Client.getTest(id);
          const coursesData = await Client.getCourses();
          setTest(testData);
          console.log("Test data:", testData);
          setCourses(coursesData);
        } catch (error) {
          console.error("An error occurred while fetching tests:", error);
        }
      }
    };

    fetchData();
  }, [id]);

  const updateTest = (courseId?: string) => {
    if (test?.id && (courseId || test.course?.id)) {
      return Client.putTest(test.id, {
        name: test.name ?? "",
        courseId: courseId ?? test.course?.id ?? "",
        schoolYear: test.schoolYear ?? "",
      });
    }
  };

  const handleTestNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTest({ ...test, name: event.target.value });
  };

  const handleSelectedCourseChange = (course: ICourse | undefined) => {
    if (course?.id !== test?.course?.id && course?.id) {
      updateTest(course?.id)?.then(() => {
        setTest({ ...test, course: course });
      });
    }
  };

  const handleTestYearChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTest({ ...test, schoolYear: event.target.value });
  };

  const handleQuestionUpdated = (question: IQuestion) => {
    setTest({
      ...test,
      questions: test?.questions?.map((q) =>
        q.id === question.id ? question : q
      ),
    });
  };

  const handleQuestionDeleted = (questionId: string) => {
    setTest({
      ...test,
      questions: test?.questions?.filter((q) => q.id !== questionId),
    });
  };

  const addNewQuestion = () => {
    if (test?.id) {
      if (test.questions?.find((q) => !q.body)) {
        toast({
          title: "Uzupełnij puste pytanie przed dodaniem nowego",
          variant: "destructive",
        });
        return;
      }
      const newQuestion = {
        testId: test.id,
        body: "",
        imgFile: "",
        answers: [],
      };
      Client.postQuestion(newQuestion).then((response) => {
        setTest({
          ...test,
          questions: [
            ...(test.questions ?? []),
            { ...newQuestion, id: response.id },
          ],
        });
      });
    }
  };
  return (
    <div className="flex min-h-screen w-full flex-col">
      <Navbar />
      <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
        <div className="grid gap-2 text-center">
          <div className="text-4xl">Edytuj testownik</div>
        </div>
        <div>
          <Label>Nazwa testownika</Label>
          <Input
            value={test?.name ?? ""}
            onChange={handleTestNameChange}
            required
            onBlur={() => updateTest()}
          />
        </div>
        <CourseSelector
          courses={courses}
          selectedCourse={test?.course}
          setSelectedCourse={handleSelectedCourseChange}
          dropdownWidth="1200px"
        />
        <div>
          <Label>Semestr</Label>
          <Input
            value={test?.schoolYear ?? ""}
            onChange={handleTestYearChange}
            onBlur={() => updateTest()}
          />
        </div>
        <div>
          <Label>Pytania</Label>
          <div className="grid gap-4">
            {test?.questions?.map((question, index) => (
              <EditQuestion
                key={question.id}
                index={index}
                question={question}
                onUpdated={handleQuestionUpdated}
                onDeleted={handleQuestionDeleted}
              />
            ))}
          </div>
        </div>
        <Button className="gap-2" onClick={addNewQuestion}>
          <Plus className="h-5 w-5" /> Dodaj nowe pytanie
        </Button>
      </main>
    </div>
  );
};

export default EditTest;
