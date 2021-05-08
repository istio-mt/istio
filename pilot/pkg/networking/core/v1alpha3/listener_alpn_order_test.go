package v1alpha3

import (
	"reflect"
	"testing"

	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	tlsv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
)

func TestOrderListenerDownstreamTlsContextAlpnProtocols(t *testing.T) {
	for _, c := range []struct {
		originAlpnProtocols []string
		wantedAlpnProtocols []string
	}{
		{
			originAlpnProtocols: []string{"http/1.1"},
			wantedAlpnProtocols: []string{"http/1.1"},
		},
		{
			originAlpnProtocols: []string{"http/1.1", "h2"},
			wantedAlpnProtocols: []string{"h2", "http/1.1"},
		},
		{
			originAlpnProtocols: []string{"http/1.1", "h2", "http/1.0"},
			wantedAlpnProtocols: []string{"h2", "http/1.1", "http/1.0"},
		},
	} {
		tlsContext := &tlsv3.DownstreamTlsContext{
			CommonTlsContext: &tlsv3.CommonTlsContext{
				AlpnProtocols: c.originAlpnProtocols,
			},
		}
		transportSocket := buildDownstreamTLSTransportSocket(tlsContext)
		listener := &listenerv3.Listener{
			FilterChains: []*listenerv3.FilterChain{
				{
					TransportSocket: transportSocket,
				},
			},
		}
		orderListenerDownstreamTlsContextAlpnProtocols([]*listenerv3.Listener{listener})

		var gotTLSContext tlsv3.DownstreamTlsContext
		_ = listener.FilterChains[0].TransportSocket.GetTypedConfig().UnmarshalTo(&gotTLSContext)
		gotAlpnProtocols := gotTLSContext.CommonTlsContext.AlpnProtocols
		if !reflect.DeepEqual(gotAlpnProtocols, c.wantedAlpnProtocols) {
			t.Errorf("orderListenerDownstreamTlsContextAlpnProtocols failed for %v:got=%v,wanted=%v", c.originAlpnProtocols, gotAlpnProtocols, c.wantedAlpnProtocols)
		}
	}
}
