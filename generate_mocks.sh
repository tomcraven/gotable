# The command to generate mocks is pretty precise
# Add any other mockgen commands to this file
mockgen -source=output.go -destination=gotable_mock/mock_output.go -imports=".=github.com/tomcraven/gotable"
