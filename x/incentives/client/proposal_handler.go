package client

import (
	govclient "cosmossdk.io/x/gov/client"

	"github.com/furyanrasta/blackfury/v3/x/incentives/client/cli"
	"github.com/furyanrasta/blackfury/v3/x/incentives/client/rest"
)

var (
	RegisterIncentiveProposalHandler = govclient.NewProposalHandler(cli.NewRegisterIncentiveProposalCmd, rest.RegisterIncentiveProposalRESTHandler)
	CancelIncentiveProposalHandler   = govclient.NewProposalHandler(cli.NewCancelIncentiveProposalCmd, rest.CancelIncentiveProposalRequestRESTHandler)
)
