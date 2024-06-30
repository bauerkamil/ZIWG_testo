namespace USOSConnector.Functions.Triggers.Dtos;

public record PagedResponseDto<T>
{
    public required int PageNumber { get; init; }
    public required int PageSize { get; init; }
    public required bool IsNextPage { get; init; }
    public required IEnumerable<T> Items { get; init; }
}