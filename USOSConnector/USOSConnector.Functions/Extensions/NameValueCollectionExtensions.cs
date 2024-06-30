using System.Collections.Specialized;
using USOSConnector.Functions.Services.UsosService.Dtos;

namespace USOSConnector.Functions.Extensions;

public static class NameValueCollectionExtensions
{
    public static PaginationDto GetPagination(this NameValueCollection collection)
    {
        var allParams = collection.AllKeys.ToDictionary(k => k!.ToLower(), k => collection[k]);

        var pagedRequest = new PaginationDto
        {
            PageNumber = 1,
            PageSize = 5
        };

        if (allParams.ContainsKey("pagenumber") && int.TryParse(allParams["pagenumber"], out var pageNumber))
        {
            pagedRequest.PageNumber = pageNumber;
        }

        if (allParams.ContainsKey("pagesize") && int.TryParse(allParams["pagesize"], out var pageSize))
        {
            pagedRequest.PageSize = pageSize;
        }

        return pagedRequest;
    }
}