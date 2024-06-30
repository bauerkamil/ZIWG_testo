using System.Text.Json.Serialization;

namespace USOSConnector.Functions.Services.UsosService.Dtos;

public record SearchCoursesDto
{
    [JsonPropertyName("items")]
    public required SearchCoursesCource[] Items { get; init; }

    [JsonPropertyName("next_page")]
    public required bool NextPage { get; init; }
}

public record SearchCoursesCource
{
    [JsonPropertyName("course_id")]
    public required string CourseId { get; init; }

    [JsonPropertyName("name")]
    public required SearchCoursesName Name { get; init; }
}

public record SearchCoursesName
{
    [JsonPropertyName("pl")]
    public required string Pl { get; init; }

    [JsonPropertyName("en")]
    public required string En { get; init; }
}