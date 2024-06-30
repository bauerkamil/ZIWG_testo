using System.Net;
using Microsoft.Azure.Functions.Worker;
using Microsoft.Azure.Functions.Worker.Http;
using Microsoft.Extensions.Caching.Memory;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using USOSConnector.Functions.Attributes;
using USOSConnector.Functions.Extensions;
using USOSConnector.Functions.Options;
using USOSConnector.Functions.Services.JwtService;
using USOSConnector.Functions.Services.UsosService;
using USOSConnector.Functions.Services.UsosService.Dtos;
using USOSConnector.Functions.Triggers.Dtos;

namespace USOSConnector.Functions.Triggers;

public class SearchCoursesTrigger
{
    private readonly IMemoryCache _cache;
    private readonly IUsosService _usosService;
    private readonly IJwtService _jwtService;
    private readonly TimeProvider _timeProvider;
    private readonly USOSOptions _options;
    private readonly ILogger<SearchCoursesTrigger> _logger;

    public SearchCoursesTrigger(
        IMemoryCache cache,
        IUsosService usosService,
        IJwtService jwtService,
        TimeProvider timeProvider,
        ILogger<SearchCoursesTrigger> logger,
        IOptions<USOSOptions> options)
    {
        _cache = cache;
        _usosService = usosService;
        _jwtService = jwtService;
        _timeProvider = timeProvider;
        _options = options.Value;
        _logger = logger;
    }

    [JwtAuthorize]
    [Function(nameof(CoursesSearch))]
    public async Task<HttpResponseData> CoursesSearch(
        [HttpTrigger(AuthorizationLevel.Anonymous, "get", Route = "courses/search")] 
        HttpRequestData req,
        CancellationToken cancellationToken)
    {
        var bearerToken = req.GetBearerToken();

        var userClaims = _jwtService.GetUserClaims(bearerToken);

        var tokenSecret = _cache.Get<string>(userClaims.UserId);

        if (string.IsNullOrEmpty(tokenSecret))
        {
            // Token is cached as long as it is valid
            return await req.CreateProblemResponseAsync(
                HttpStatusCode.Unauthorized,
                "Expired user token.");
        }

        var searchQuery = req.Query["query"];

        if (string.IsNullOrEmpty(searchQuery))
        {
            return await req.CreateProblemResponseAsync(
                HttpStatusCode.BadRequest,
                "'query' parameter is required.");
        }

        var pagination = req.Query.GetPagination();

        var searchResponse = await _usosService.SearchForCoursesAsync(
            userClaims.UsosToken,
            tokenSecret,
            searchQuery,
            pagination,
            cancellationToken);

        var pagedResult = MapToResponseDto(pagination, searchResponse);

        return await req.CreateOkResponseAsync(pagedResult);
    }

    private static PagedResponseDto<SearchCoursesResponseDto> MapToResponseDto(
        PaginationDto pagination, 
        SearchCoursesDto searchResponse) => new PagedResponseDto<SearchCoursesResponseDto>
        {
            PageNumber = pagination.PageNumber,
            PageSize = pagination.PageSize,
            IsNextPage = searchResponse.NextPage,
            Items = searchResponse.Items
                .Select(i => new SearchCoursesResponseDto
                {
                    CourseId = i.CourseId,
                    CourseName = i.Name.Pl
                })
        };
}