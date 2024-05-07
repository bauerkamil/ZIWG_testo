import React from "react";
import {
  Button,
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui";
import { EllipsisVertical, PencilLine, Trash } from "lucide-react";
import { useNavigate } from "react-router-dom";

const TestCardOptions = (props: { id: string }) => {
  const { id } = props;
  const navigate = useNavigate();

  const editTest = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.stopPropagation();
    navigate(`/edit/${id}`);
  };

  const deleteTest = (event: React.MouseEvent<HTMLButtonElement>) => {
    event.stopPropagation();
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" size="icon" className="rounded-full">
          <EllipsisVertical className="h-4 w-4" />
          <span className="sr-only">Toggle user menu</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuItem>
          <Button className="h-5" variant="ghost" onClick={editTest}>
            <div className="flex flex-row gap-2 align-center justify-center">
              <PencilLine className="h-4 w-4" />
              Edytuj
            </div>
          </Button>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Button className="h-5" variant="ghost" onClick={deleteTest}>
            <div className="flex flex-row gap-2 align-center justify-center">
              <Trash className="h-4 w-4" />
              Usuń
            </div>
          </Button>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default TestCardOptions;
