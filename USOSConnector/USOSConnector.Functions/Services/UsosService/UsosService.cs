using System.Net.Http.Json;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using USOSConnector.Functions.Constants;
using USOSConnector.Functions.Dtos;
using USOSConnector.Functions.Helpers;
using USOSConnector.Functions.Options;

namespace USOSConnector.Functions.Services.UsosService;

public class UsosService : IUsosService
{
    private readonly USOSOptions _options;
    private readonly HttpClient _httpClient;
    private readonly TimeProvider _timeProvider;

    public UsosService(
        HttpClient httpClient,
        TimeProvider timeProvider,
        IOptions<USOSOptions> options)
    {
        _options = options.Value;
        _httpClient = httpClient;
        _timeProvider = timeProvider;
    }

    public async Task<RequestTokenDto> GetRequestTokenAsync(
        string? callbackKey = null, 
        string callbackKeyName = "key",
        CancellationToken cancellationToken = default)
    {
        var (consumerKey, consumerSecret, callbackUrl, apiUrl) = _options;

        if (callbackKey is not null)
        {
            callbackUrl = callbackUrl + $"?{callbackKeyName}={callbackKey}";
        }

        var query = new Dictionary<string, string>
        {
            ["oauth_callback"] = Uri.EscapeDataString(callbackUrl),
        };

        var coursesResponse = await CallEndpointAsync(
            Usos.Endpoints.RequestToken,
            query,
            string.Empty, 
            cancellationToken);

        var result = await coursesResponse.Content.ReadAsStringAsync(cancellationToken);

        var parts = result.Split('&')
            .Select(x => x.Split('='))
            .ToDictionary(x => x[0], x => x[1]);

        var requestTokenResult = new RequestTokenDto
        {
            Secret = parts["oauth_token_secret"],
            RedirectUri = $"{_options.ApiUrl}{Usos.Endpoints.Authorize}?oauth_token={parts["oauth_token"]}"
        };
        return requestTokenResult;
    }

    public async Task<AccessTokenDto> GetAccessTokenAsync(
        string token, 
        string verifier, 
        string secret, 
        CancellationToken cancellationToken)
    {
        var query = new Dictionary<string, string>
        {
            ["oauth_token"] = token,
            ["oauth_verifier"] = verifier,
        };

        var coursesResponse = await CallEndpointAsync(
            Usos.Endpoints.AccessToken,
            query,
            secret, 
            cancellationToken);

        var result = await coursesResponse.Content.ReadAsStringAsync(cancellationToken);

        var parts = result.Split('&')
            .Select(x => x.Split('='))
            .ToDictionary(x => x[0], x => x[1]);

        var accessTokenResult = new AccessTokenDto
        {
            Token = parts["oauth_token"],
            Secret = parts["oauth_token_secret"]
        };
        return accessTokenResult;
    }

    public async Task<string> GetUserCoursesAsync(
        string secret, 
        string token, 
        CancellationToken cancellationToken)
    {
        var query = new Dictionary<string, string>
        {
            ["oauth_token"] = token,
        };

        var coursesResponse = await CallEndpointAsync(
            Usos.Endpoints.UserCourses,
            query,
            secret, 
            cancellationToken);

        var userCoursesResult = await coursesResponse.Content.ReadAsStringAsync(cancellationToken);
        return userCoursesResult;
    }

    public async Task<UserInfoDto> GetCurrentUserAsync(
        string token, 
        string secret, 
        CancellationToken cancellationToken)
    {
        var query = new Dictionary<string, string>
        {
            ["oauth_token"] = token,
        };

        var coursesResponse = await CallEndpointAsync(
            Usos.Endpoints.UserInfo,
            query, 
            secret, 
            cancellationToken);

        var userInfoResult = await coursesResponse.Content.ReadFromJsonAsync<UserInfoDto>(cancellationToken);
        ArgumentNullException.ThrowIfNull(userInfoResult);
        return userInfoResult;
    }

    private async Task<HttpResponseMessage> CallEndpointAsync(
        string endpoint,
        Dictionary<string, string> query, 
        string secret, 
        CancellationToken cancellationToken)
    {
        var (consumerKey, consumerSecret, _, apiUrl) = _options;

        query["oauth_consumer_key"] = consumerKey;
        query["oauth_nonce"] = Guid.NewGuid().ToString();
        query["oauth_timestamp"] = _timeProvider.GetUtcNow().ToUnixTimeSeconds().ToString();
        query["oauth_signature_method"] = "HMAC-SHA1";
        query["oauth_version"] = "1.0";

        var fullEndpoint = apiUrl + endpoint;
        var key = consumerSecret + "&" + secret;
        string coursesUri = OAuthHelper.GetUri(fullEndpoint, key, query);

        var response = await _httpClient.GetAsync(coursesUri, cancellationToken);
        return response;
    }
}