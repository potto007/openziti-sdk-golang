/*
	Copyright 2019 NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package ziti

// Deprecated: DefaultCollection is deprecated and is included for legacy support.
// It powers two other deprecated functions: `ForAllContext() and and `LoadContext()` which rely on it. The intended
// replacement is for implementations that wish to have this functionality to use NewSdkCollection() or
// NewSdkCollectionFromEnv() on their own.
var DefaultCollection *SdkCollection

func init() {
	DefaultCollection = NewSdkCollectionFromEnv("ZITI_IDENTITIES")
	DefaultCollection.ConfigTypes = []string{InterceptV1, ClientConfigV1}
}

// Deprecated: ForAllContexts iterates over all Context instances in the DefaultCollection and call the provided function `f`.
// Usage of the DefaultCollection is advised against, and if this functionality is needed, implementations should
// instantiate their own SdkCollection via NewSdkCollection() or NewSdkCollectionFromEnv()
func ForAllContexts(f func(ctx Context) bool) {
	DefaultCollection.ForAll(f)
}

// Deprecated: LoadContext loads a configuration from the supplied path into the DefaultCollection as a convenience.
// Usage of the DefaultCollection is advised against, and if this functionality is needed, implementations should
// instantiate their own SdkCollection via NewSdkCollection() or NewSdkCollectionFromEnv().
//
// This function's behavior can be replicated with:
// ```
//
// collection = NewSdkCollection()
// collection.ConfigTypes = []string{InterceptV1, ClientConfigV1}
// collection.NewContextFromFile(configPath)
//
// ```
//
// LoadContext will attempt to load a Config from the provided path, see NewConfigFromFile() for details. Additionally,
// LoadContext will attempt to authenticate the Context. If it does not authenticate, it will not be added to the
// DefaultCollection and an error will be returned.
// ```
func LoadContext(configPath string) (Context, error) {
	ctx, err := DefaultCollection.NewContextFromFile(configPath)

	if err != nil {
		return nil, err
	}

	err = ctx.Authenticate()

	if err != nil {
		DefaultCollection.Remove(ctx)
		ctx.Close()
	}

	return ctx, nil
}
