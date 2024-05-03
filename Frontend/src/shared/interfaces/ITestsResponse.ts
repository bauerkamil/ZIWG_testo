export interface ITestsResponse {
    id: string;
    name: string;
    createdBy: string;
    courseId: string;
    createdAt: string;
    changedBy: string | null;
    changedAt: string | null;
    course: {
        id: string;
        name: string;
        teacher_id: string;
        teacher: {
            id: string;
            name: string;
            second_name: string;
            surname: string;
        };
        school_year: number;
    };
}