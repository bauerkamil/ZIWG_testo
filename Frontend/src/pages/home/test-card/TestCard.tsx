import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card";
import { ITestInfo } from "@/shared/interfaces";

const TestCard= (props: {test: ITestInfo}) => {
  const {test} = props;

  return (
    <Card className="cursor-pointer hover:border-solid hover:border-current group">
      <CardHeader>
        <div className="font-semibold leading-none tracking-tight group-hover:text-2xl">
          {test.name}
        </div>
      </CardHeader>
      <CardContent>        
        <div className="text-primary font-bold">{test.course}</div>
        <div className="flex flex-row">
          <p className="text-muted-foreground italic">prowadzący:&nbsp;</p> 
          <p>{test.teacher}</p>
        </div>
      </CardContent>
      <CardFooter>
        <div className="text-sm text-muted-foreground">Ostatnia modyfikacja: {test.lastModified.toDateString()}</div>
      </CardFooter>
    </Card>

  );
};

export default TestCard;
