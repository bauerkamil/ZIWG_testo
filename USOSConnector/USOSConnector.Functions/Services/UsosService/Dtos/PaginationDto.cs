namespace USOSConnector.Functions.Services.UsosService.Dtos;

public record PaginationDto
{
    public int PageNumber { get; set; }
    public int PageSize { get; set; }
}