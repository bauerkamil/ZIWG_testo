using Microsoft.Extensions.Options;
using Microsoft.IdentityModel.Tokens;
using USOSConnector.Functions.Options;
using USOSConnector.Functions.Services.JwtService;
using USOSConnector.Functions.Services.JwtService.Dtos;

namespace USOSConnector.Tests.UnitTests;

public class JwtServiceTests
{
    [Fact]
    public void IsTokenValid_WithValidToken_ShouldReturnTrue()
    {
        // Arrange
        var options = Options.Create(new JwtOptions
        {
            Key = Guid.NewGuid().ToString(),
            ExpiryMinutes = 1
        });

        var jwtService = new JwtService(options);

        var token = jwtService.GenerateToken(new UserClaimsDto
        {
            UserId = "userId",
            FirstName = "firstName",
            LastName = "lastName",
            UsosToken = "usosToken"
        });

        // Act
        var actual = jwtService.IsTokenValid(token);

        // Assert
        Assert.True(actual);
    }

    [Fact]
    public void IsTokenValid_WithInvalidToken_ShouldReturnFalse()
    {
        // Arrange
        var options = Options.Create(new JwtOptions
        {
            Key = Guid.NewGuid().ToString(),
            ExpiryMinutes = 1
        });

        var jwtService = new JwtService(options);

        // Act
        var actual = jwtService.IsTokenValid("invalidToken");

        // Assert
        Assert.False(actual);
    }

    // UserUserClaims invalid token throws exception
    [Fact]
    public void GetUserClaims_WithInvalidToken_ShouldThrowException()
    {
        // Arrange
        var options = Options.Create(new JwtOptions
        {
            Key = Guid.NewGuid().ToString(),
            ExpiryMinutes = 1
        });

        var jwtService = new JwtService(options);

        // Act 
        void Act() => jwtService.GetUserClaims("invalidToken");

        //Assert
        Assert.Throws<SecurityTokenMalformedException>(Act);
    }
}