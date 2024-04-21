package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/NidroidX/kestrelcoind/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.kestrelcoindMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.kestrelcoindMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.kestrelcoindMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.kestrelcoindMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.kestrelcoindMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.kestrelcoindMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.kestrelcoindMessage_BanRequest{}),
	reflect.TypeOf(protowire.kestrelcoindMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
