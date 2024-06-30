namespace USOSConnector.Functions.Triggers.Dtos;

public record SearchCoursesResponseDto
{
    public required string CourseId { get; init; }
    public required string CourseName { get; init; }
}