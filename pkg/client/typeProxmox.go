package client

type Cluster struct {
	Children []struct {
		Children []struct {
			Info struct {
				DELETE struct {
					Allowtoken  int    `json:"allowtoken"`
					Description string `json:"description"`
					Method      string `json:"method"`
					Name        string `json:"name"`
					Parameters  struct {
						AdditionalProperties int `json:"additionalProperties"`
						Properties           struct {
							Force struct {
								Default     int    `json:"default"`
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"force"`
							ID struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								Pattern     string `json:"pattern"`
								Type        string `json:"type"`
							} `json:"id"`
							Keep struct {
								Default     int    `json:"default"`
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"keep"`
						} `json:"properties"`
					} `json:"parameters"`
					Permissions struct {
						Check []interface{} `json:"check"`
					} `json:"permissions"`
					Protected int `json:"protected"`
					Returns   struct {
						Type string `json:"type"`
					} `json:"returns"`
				} `json:"DELETE"`
				GET struct {
					Allowtoken  int    `json:"allowtoken"`
					Description string `json:"description"`
					Method      string `json:"method"`
					Name        string `json:"name"`
					Parameters  struct {
						AdditionalProperties int `json:"additionalProperties"`
						Properties           struct {
							ID struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								Pattern     string `json:"pattern"`
								Type        string `json:"type"`
							} `json:"id"`
						} `json:"properties"`
					} `json:"parameters"`
					Permissions struct {
						Description string `json:"description"`
						User        string `json:"user"`
					} `json:"permissions"`
					Returns struct {
						Type string `json:"type"`
					} `json:"returns"`
				} `json:"GET"`
				PUT struct {
					Allowtoken  int    `json:"allowtoken"`
					Description string `json:"description"`
					Method      string `json:"method"`
					Name        string `json:"name"`
					Parameters  struct {
						AdditionalProperties int `json:"additionalProperties"`
						Properties           struct {
							Comment struct {
								Description string `json:"description"`
								MaxLength   int    `json:"maxLength"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"comment"`
							Delete struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								MaxLength   int    `json:"maxLength"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"delete"`
							Digest struct {
								Description string `json:"description"`
								MaxLength   int    `json:"maxLength"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"digest"`
							Disable struct {
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"disable"`
							ID struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								Pattern     string `json:"pattern"`
								Type        string `json:"type"`
							} `json:"id"`
							Rate struct {
								Description string `json:"description"`
								Minimum     int    `json:"minimum"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"rate"`
							RemoveJob struct {
								Description string   `json:"description"`
								Enum        []string `json:"enum"`
								Optional    int      `json:"optional"`
								Type        string   `json:"type"`
							} `json:"remove_job"`
							Schedule struct {
								Default     string `json:"default"`
								Description string `json:"description"`
								Format      string `json:"format"`
								MaxLength   int    `json:"maxLength"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"schedule"`
							Source struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"source"`
						} `json:"properties"`
						Type string `json:"type"`
					} `json:"parameters"`
					Permissions struct {
						Check []interface{} `json:"check"`
					} `json:"permissions"`
					Protected int `json:"protected"`
					Returns   struct {
						Type string `json:"type"`
					} `json:"returns"`
				} `json:"PUT"`
			} `json:"info"`
			Leaf int    `json:"leaf"`
			Path string `json:"path"`
			Text string `json:"text"`
		} `json:"children,omitempty"`
		Info struct {
			GET struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
				} `json:"parameters"`
				Permissions struct {
					Description string `json:"description"`
					User        string `json:"user"`
				} `json:"permissions"`
				Returns struct {
					Items struct {
						Properties struct {
						} `json:"properties"`
						Type string `json:"type"`
					} `json:"items"`
					Links []struct {
						Href string `json:"href"`
						Rel  string `json:"rel"`
					} `json:"links"`
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"GET"`
			POST struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Comment struct {
							Description string `json:"description"`
							MaxLength   int    `json:"maxLength"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"comment"`
						Disable struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"disable"`
						ID struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Pattern     string `json:"pattern"`
							Type        string `json:"type"`
						} `json:"id"`
						Rate struct {
							Description string `json:"description"`
							Minimum     int    `json:"minimum"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"rate"`
						RemoveJob struct {
							Description string   `json:"description"`
							Enum        []string `json:"enum"`
							Optional    int      `json:"optional"`
							Type        string   `json:"type"`
						} `json:"remove_job"`
						Schedule struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							MaxLength   int    `json:"maxLength"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"schedule"`
						Source struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"source"`
						Target struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"target"`
						Type struct {
							Description string   `json:"description"`
							Enum        []string `json:"enum"`
							Type        string   `json:"type"`
						} `json:"type"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"parameters"`
				Permissions struct {
					Check []interface{} `json:"check"`
				} `json:"permissions"`
				Protected int `json:"protected"`
				Returns   struct {
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"POST"`
		} `json:"info,omitempty"`
		Leaf int    `json:"leaf"`
		Path string `json:"path"`
		Text string `json:"text"`
	} `json:"children"`
	Info struct {
		GET struct {
			Allowtoken  int    `json:"allowtoken"`
			Description string `json:"description"`
			Method      string `json:"method"`
			Name        string `json:"name"`
			Parameters  struct {
				AdditionalProperties int `json:"additionalProperties"`
			} `json:"parameters"`
			Permissions struct {
				User string `json:"user"`
			} `json:"permissions"`
			Returns struct {
				Items struct {
					Properties struct {
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"items"`
				Links []struct {
					Href string `json:"href"`
					Rel  string `json:"rel"`
				} `json:"links"`
				Type string `json:"type"`
			} `json:"returns"`
		} `json:"GET"`
	} `json:"info"`
	Leaf int    `json:"leaf"`
	Path string `json:"path"`
	Text string `json:"text"`
}

type Nodes struct {
	Children []struct {
		Children []struct {
			Children []struct {
				Children []struct {
					Children []struct {
						Children []struct {
							Info struct {
								DELETE struct {
									Allowtoken  int    `json:"allowtoken"`
									Description string `json:"description"`
									Method      string `json:"method"`
									Name        string `json:"name"`
									Parameters  struct {
										AdditionalProperties int `json:"additionalProperties"`
										Properties           struct {
											Digest struct {
												Description string `json:"description"`
												MaxLength   int    `json:"maxLength"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"digest"`
											Node struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"node"`
											Pos struct {
												Description string `json:"description"`
												Minimum     int    `json:"minimum"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"pos"`
											Vmid struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Minimum     int    `json:"minimum"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"vmid"`
										} `json:"properties"`
									} `json:"parameters"`
									Permissions struct {
										Check []interface{} `json:"check"`
									} `json:"permissions"`
									Protected int         `json:"protected"`
									Proxyto   interface{} `json:"proxyto"`
									Returns   struct {
										Type string `json:"type"`
									} `json:"returns"`
								} `json:"DELETE"`
								GET struct {
									Allowtoken  int    `json:"allowtoken"`
									Description string `json:"description"`
									Method      string `json:"method"`
									Name        string `json:"name"`
									Parameters  struct {
										AdditionalProperties int `json:"additionalProperties"`
										Properties           struct {
											Node struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"node"`
											Pos struct {
												Description string `json:"description"`
												Minimum     int    `json:"minimum"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"pos"`
											Vmid struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Minimum     int    `json:"minimum"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"vmid"`
										} `json:"properties"`
									} `json:"parameters"`
									Permissions struct {
										Check []interface{} `json:"check"`
									} `json:"permissions"`
									Proxyto interface{} `json:"proxyto"`
									Returns struct {
										Properties struct {
											Action struct {
												Type string `json:"type"`
											} `json:"action"`
											Comment struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"comment"`
											Dest struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"dest"`
											Dport struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"dport"`
											Enable struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"enable"`
											Iface struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"iface"`
											Ipversion struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"ipversion"`
											Log struct {
												Description string   `json:"description"`
												Enum        []string `json:"enum"`
												Optional    int      `json:"optional"`
												Type        string   `json:"type"`
											} `json:"log"`
											Macro struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"macro"`
											Pos struct {
												Type string `json:"type"`
											} `json:"pos"`
											Proto struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"proto"`
											Source struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"source"`
											Sport struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"sport"`
											Type struct {
												Type string `json:"type"`
											} `json:"type"`
										} `json:"properties"`
										Type string `json:"type"`
									} `json:"returns"`
								} `json:"GET"`
								PUT struct {
									Allowtoken  int    `json:"allowtoken"`
									Description string `json:"description"`
									Method      string `json:"method"`
									Name        string `json:"name"`
									Parameters  struct {
										AdditionalProperties int `json:"additionalProperties"`
										Properties           struct {
											Action struct {
												Description string `json:"description"`
												MaxLength   int    `json:"maxLength"`
												MinLength   int    `json:"minLength"`
												Optional    int    `json:"optional"`
												Pattern     string `json:"pattern"`
												Type        string `json:"type"`
											} `json:"action"`
											Comment struct {
												Description string `json:"description"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"comment"`
											Delete struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"delete"`
											Dest struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"dest"`
											Digest struct {
												Description string `json:"description"`
												MaxLength   int    `json:"maxLength"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"digest"`
											Dport struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"dport"`
											Enable struct {
												Description string `json:"description"`
												Minimum     int    `json:"minimum"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"enable"`
											Iface struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												MaxLength   int    `json:"maxLength"`
												MinLength   int    `json:"minLength"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"iface"`
											Log struct {
												Description string   `json:"description"`
												Enum        []string `json:"enum"`
												Optional    int      `json:"optional"`
												Type        string   `json:"type"`
											} `json:"log"`
											Macro struct {
												Description string `json:"description"`
												MaxLength   int    `json:"maxLength"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"macro"`
											Moveto struct {
												Description string `json:"description"`
												Minimum     int    `json:"minimum"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"moveto"`
											Node struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"node"`
											Pos struct {
												Description string `json:"description"`
												Minimum     int    `json:"minimum"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"pos"`
											Proto struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"proto"`
											Source struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"source"`
											Sport struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"sport"`
											Type struct {
												Description string   `json:"description"`
												Enum        []string `json:"enum"`
												Optional    int      `json:"optional"`
												Type        string   `json:"type"`
											} `json:"type"`
											Vmid struct {
												Description string `json:"description"`
												Format      string `json:"format"`
												Minimum     int    `json:"minimum"`
												Type        string `json:"type"`
												Typetext    string `json:"typetext"`
											} `json:"vmid"`
										} `json:"properties"`
									} `json:"parameters"`
									Permissions struct {
										Check []interface{} `json:"check"`
									} `json:"permissions"`
									Protected int         `json:"protected"`
									Proxyto   interface{} `json:"proxyto"`
									Returns   struct {
										Type string `json:"type"`
									} `json:"returns"`
								} `json:"PUT"`
							} `json:"info"`
							Leaf int    `json:"leaf"`
							Path string `json:"path"`
							Text string `json:"text"`
						} `json:"children,omitempty"`
						Leaf int    `json:"leaf"`
						Path string `json:"path"`
						Text string `json:"text"`
					} `json:"children,omitempty"`
					Info struct {
						GET struct {
							Allowtoken  int    `json:"allowtoken"`
							Description string `json:"description"`
							Method      string `json:"method"`
							Name        string `json:"name"`
							Parameters  struct {
								AdditionalProperties int `json:"additionalProperties"`
								Properties           struct {
									Node struct {
										Description string `json:"description"`
										Format      string `json:"format"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"node"`
									Vmid struct {
										Description string `json:"description"`
										Format      string `json:"format"`
										Minimum     int    `json:"minimum"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"vmid"`
								} `json:"properties"`
							} `json:"parameters"`
							Permissions struct {
								User string `json:"user"`
							} `json:"permissions"`
							Returns struct {
								Items struct {
									Properties struct {
									} `json:"properties"`
									Type string `json:"type"`
								} `json:"items"`
								Links []struct {
									Href string `json:"href"`
									Rel  string `json:"rel"`
								} `json:"links"`
								Type string `json:"type"`
							} `json:"returns"`
						} `json:"GET"`
					} `json:"info,omitempty"`
					Leaf int    `json:"leaf"`
					Path string `json:"path"`
					Text string `json:"text"`
				} `json:"children"`
				Info struct {
					DELETE struct {
						Allowtoken  int    `json:"allowtoken"`
						Description string `json:"description"`
						Method      string `json:"method"`
						Name        string `json:"name"`
						Parameters  struct {
							AdditionalProperties int `json:"additionalProperties"`
							Properties           struct {
								Node struct {
									Description string `json:"description"`
									Format      string `json:"format"`
									Type        string `json:"type"`
									Typetext    string `json:"typetext"`
								} `json:"node"`
								Purge struct {
									Description string `json:"description"`
									Optional    int    `json:"optional"`
									Type        string `json:"type"`
									Typetext    string `json:"typetext"`
								} `json:"purge"`
								Skiplock struct {
									Description string `json:"description"`
									Optional    int    `json:"optional"`
									Type        string `json:"type"`
									Typetext    string `json:"typetext"`
								} `json:"skiplock"`
								Vmid struct {
									Description string `json:"description"`
									Format      string `json:"format"`
									Minimum     int    `json:"minimum"`
									Type        string `json:"type"`
									Typetext    string `json:"typetext"`
								} `json:"vmid"`
							} `json:"properties"`
						} `json:"parameters"`
						Permissions struct {
							Check []interface{} `json:"check"`
						} `json:"permissions"`
						Protected int    `json:"protected"`
						Proxyto   string `json:"proxyto"`
						Returns   struct {
							Type string `json:"type"`
						} `json:"returns"`
					} `json:"DELETE"`
					GET struct {
						Allowtoken  int    `json:"allowtoken"`
						Description string `json:"description"`
						Method      string `json:"method"`
						Name        string `json:"name"`
						Parameters  struct {
							AdditionalProperties int `json:"additionalProperties"`
							Properties           struct {
								Node struct {
									Description string `json:"description"`
									Format      string `json:"format"`
									Type        string `json:"type"`
									Typetext    string `json:"typetext"`
								} `json:"node"`
								Vmid struct {
									Description string `json:"description"`
									Format      string `json:"format"`
									Minimum     int    `json:"minimum"`
									Type        string `json:"type"`
									Typetext    string `json:"typetext"`
								} `json:"vmid"`
							} `json:"properties"`
						} `json:"parameters"`
						Permissions struct {
							User string `json:"user"`
						} `json:"permissions"`
						Proxyto string `json:"proxyto"`
						Returns struct {
							Items struct {
								Properties struct {
									Subdir struct {
										Type string `json:"type"`
									} `json:"subdir"`
								} `json:"properties"`
								Type string `json:"type"`
							} `json:"items"`
							Links []struct {
								Href string `json:"href"`
								Rel  string `json:"rel"`
							} `json:"links"`
							Type string `json:"type"`
						} `json:"returns"`
					} `json:"GET"`
				} `json:"info"`
				Leaf int    `json:"leaf"`
				Path string `json:"path"`
				Text string `json:"text"`
			} `json:"children,omitempty"`
			Leaf int    `json:"leaf"`
			Path string `json:"path"`
			Text string `json:"text"`
		} `json:"children"`
		Info struct {
			GET struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Node struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"node"`
					} `json:"properties"`
				} `json:"parameters"`
				Permissions struct {
					User string `json:"user"`
				} `json:"permissions"`
				Returns struct {
					Items struct {
						Properties struct {
						} `json:"properties"`
						Type string `json:"type"`
					} `json:"items"`
					Links []struct {
						Href string `json:"href"`
						Rel  string `json:"rel"`
					} `json:"links"`
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"GET"`
		} `json:"info"`
		Leaf int    `json:"leaf"`
		Path string `json:"path"`
		Text string `json:"text"`
	} `json:"children"`
	Info struct {
		GET struct {
			Allowtoken  int    `json:"allowtoken"`
			Description string `json:"description"`
			Method      string `json:"method"`
			Name        string `json:"name"`
			Parameters  struct {
				AdditionalProperties int `json:"additionalProperties"`
			} `json:"parameters"`
			Permissions struct {
				User string `json:"user"`
			} `json:"permissions"`
			Returns struct {
				Items struct {
					Properties struct {
						CPU struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Renderer    string `json:"renderer"`
							Type        string `json:"type"`
						} `json:"cpu"`
						Level struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
						} `json:"level"`
						Maxcpu struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
						} `json:"maxcpu"`
						Maxmem struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Renderer    string `json:"renderer"`
							Type        string `json:"type"`
						} `json:"maxmem"`
						Mem struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Renderer    string `json:"renderer"`
							Type        string `json:"type"`
						} `json:"mem"`
						Node struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Type        string `json:"type"`
						} `json:"node"`
						SslFingerprint struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
						} `json:"ssl_fingerprint"`
						Status struct {
							Description string   `json:"description"`
							Enum        []string `json:"enum"`
							Type        string   `json:"type"`
						} `json:"status"`
						Uptime struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Renderer    string `json:"renderer"`
							Type        string `json:"type"`
						} `json:"uptime"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"items"`
				Links []struct {
					Href string `json:"href"`
					Rel  string `json:"rel"`
				} `json:"links"`
				Type string `json:"type"`
			} `json:"returns"`
		} `json:"GET"`
	} `json:"info"`
	Leaf int    `json:"leaf"`
	Path string `json:"path"`
	Text string `json:"text"`
}

type Storage struct {
	Children []struct {
		Info struct {
			DELETE struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Storage struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"storage"`
					} `json:"properties"`
				} `json:"parameters"`
				Permissions struct {
					Check []interface{} `json:"check"`
				} `json:"permissions"`
				Protected int `json:"protected"`
				Returns   struct {
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"DELETE"`
			GET struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Storage struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"storage"`
					} `json:"properties"`
				} `json:"parameters"`
				Permissions struct {
					Check []interface{} `json:"check"`
				} `json:"permissions"`
				Returns struct {
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"GET"`
			PUT struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Blocksize struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"blocksize"`
						Bwlimit struct {
							Description string `json:"description"`
							Format      struct {
								Clone struct {
									Description       string `json:"description"`
									FormatDescription string `json:"format_description"`
									Minimum           string `json:"minimum"`
									Optional          int    `json:"optional"`
									Type              string `json:"type"`
								} `json:"clone"`
								Default struct {
									Description       string `json:"description"`
									FormatDescription string `json:"format_description"`
									Minimum           string `json:"minimum"`
									Optional          int    `json:"optional"`
									Type              string `json:"type"`
								} `json:"default"`
								Migration struct {
									Description       string `json:"description"`
									FormatDescription string `json:"format_description"`
									Minimum           string `json:"minimum"`
									Optional          int    `json:"optional"`
									Type              string `json:"type"`
								} `json:"migration"`
								Move struct {
									Description       string `json:"description"`
									FormatDescription string `json:"format_description"`
									Minimum           string `json:"minimum"`
									Optional          int    `json:"optional"`
									Type              string `json:"type"`
								} `json:"move"`
								Restore struct {
									Description       string `json:"description"`
									FormatDescription string `json:"format_description"`
									Minimum           string `json:"minimum"`
									Optional          int    `json:"optional"`
									Type              string `json:"type"`
								} `json:"restore"`
							} `json:"format"`
							Optional int    `json:"optional"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"bwlimit"`
						ComstarHg struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"comstar_hg"`
						ComstarTg struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"comstar_tg"`
						Content struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"content"`
						Delete struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							MaxLength   int    `json:"maxLength"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"delete"`
						Digest struct {
							Description string `json:"description"`
							MaxLength   int    `json:"maxLength"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"digest"`
						Disable struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"disable"`
						Domain struct {
							Description string `json:"description"`
							MaxLength   int    `json:"maxLength"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"domain"`
						Format struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"format"`
						Fuse struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"fuse"`
						IsMountpoint struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"is_mountpoint"`
						Krbd struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"krbd"`
						LioTpg struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"lio_tpg"`
						Maxfiles struct {
							Description string `json:"description"`
							Minimum     int    `json:"minimum"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"maxfiles"`
						Mkdir struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"mkdir"`
						Monhost struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"monhost"`
						Mountpoint struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"mountpoint"`
						Nodes struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"nodes"`
						Nowritecache struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"nowritecache"`
						Options struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"options"`
						Password struct {
							Description string `json:"description"`
							MaxLength   int    `json:"maxLength"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"password"`
						Pool struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"pool"`
						Redundancy struct {
							Default     int    `json:"default"`
							Description string `json:"description"`
							Maximum     int    `json:"maximum"`
							Minimum     int    `json:"minimum"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"redundancy"`
						Saferemove struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"saferemove"`
						SaferemoveThroughput struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"saferemove_throughput"`
						Server struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"server"`
						Server2 struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Requires    string `json:"requires"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"server2"`
						Shared struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"shared"`
						Smbversion struct {
							Description string   `json:"description"`
							Enum        []string `json:"enum"`
							Optional    int      `json:"optional"`
							Type        string   `json:"type"`
						} `json:"smbversion"`
						Sparse struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"sparse"`
						Storage struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"storage"`
						Subdir struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"subdir"`
						TaggedOnly struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"tagged_only"`
						Transport struct {
							Description string   `json:"description"`
							Enum        []string `json:"enum"`
							Optional    int      `json:"optional"`
							Type        string   `json:"type"`
						} `json:"transport"`
						Username struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"username"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"parameters"`
				Permissions struct {
					Check []interface{} `json:"check"`
				} `json:"permissions"`
				Protected int `json:"protected"`
				Returns   struct {
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"PUT"`
		} `json:"info"`
		Leaf int    `json:"leaf"`
		Path string `json:"path"`
		Text string `json:"text"`
	} `json:"children"`
	Info struct {
		GET struct {
			Allowtoken  int    `json:"allowtoken"`
			Description string `json:"description"`
			Method      string `json:"method"`
			Name        string `json:"name"`
			Parameters  struct {
				AdditionalProperties int `json:"additionalProperties"`
				Properties           struct {
					Type struct {
						Description string   `json:"description"`
						Enum        []string `json:"enum"`
						Optional    int      `json:"optional"`
						Type        string   `json:"type"`
					} `json:"type"`
				} `json:"properties"`
			} `json:"parameters"`
			Permissions struct {
				Description string `json:"description"`
				User        string `json:"user"`
			} `json:"permissions"`
			Returns struct {
				Items struct {
					Properties struct {
						Storage struct {
							Type string `json:"type"`
						} `json:"storage"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"items"`
				Links []struct {
					Href string `json:"href"`
					Rel  string `json:"rel"`
				} `json:"links"`
				Type string `json:"type"`
			} `json:"returns"`
		} `json:"GET"`
		POST struct {
			Allowtoken  int    `json:"allowtoken"`
			Description string `json:"description"`
			Method      string `json:"method"`
			Name        string `json:"name"`
			Parameters  struct {
				AdditionalProperties int `json:"additionalProperties"`
				Properties           struct {
					Authsupported struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"authsupported"`
					Base struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"base"`
					Blocksize struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"blocksize"`
					Bwlimit struct {
						Description string `json:"description"`
						Format      struct {
							Clone struct {
								Description       string `json:"description"`
								FormatDescription string `json:"format_description"`
								Minimum           string `json:"minimum"`
								Optional          int    `json:"optional"`
								Type              string `json:"type"`
							} `json:"clone"`
							Default struct {
								Description       string `json:"description"`
								FormatDescription string `json:"format_description"`
								Minimum           string `json:"minimum"`
								Optional          int    `json:"optional"`
								Type              string `json:"type"`
							} `json:"default"`
							Migration struct {
								Description       string `json:"description"`
								FormatDescription string `json:"format_description"`
								Minimum           string `json:"minimum"`
								Optional          int    `json:"optional"`
								Type              string `json:"type"`
							} `json:"migration"`
							Move struct {
								Description       string `json:"description"`
								FormatDescription string `json:"format_description"`
								Minimum           string `json:"minimum"`
								Optional          int    `json:"optional"`
								Type              string `json:"type"`
							} `json:"move"`
							Restore struct {
								Description       string `json:"description"`
								FormatDescription string `json:"format_description"`
								Minimum           string `json:"minimum"`
								Optional          int    `json:"optional"`
								Type              string `json:"type"`
							} `json:"restore"`
						} `json:"format"`
						Optional int    `json:"optional"`
						Type     string `json:"type"`
						Typetext string `json:"typetext"`
					} `json:"bwlimit"`
					ComstarHg struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"comstar_hg"`
					ComstarTg struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"comstar_tg"`
					Content struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"content"`
					Disable struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"disable"`
					Domain struct {
						Description string `json:"description"`
						MaxLength   int    `json:"maxLength"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"domain"`
					Export struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"export"`
					Format struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"format"`
					Fuse struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"fuse"`
					IsMountpoint struct {
						Default     string `json:"default"`
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"is_mountpoint"`
					Iscsiprovider struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"iscsiprovider"`
					Krbd struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"krbd"`
					LioTpg struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"lio_tpg"`
					Maxfiles struct {
						Description string `json:"description"`
						Minimum     int    `json:"minimum"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"maxfiles"`
					Mkdir struct {
						Default     string `json:"default"`
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"mkdir"`
					Monhost struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"monhost"`
					Mountpoint struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"mountpoint"`
					Nodes struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"nodes"`
					Nowritecache struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"nowritecache"`
					Options struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"options"`
					Password struct {
						Description string `json:"description"`
						MaxLength   int    `json:"maxLength"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"password"`
					Path struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"path"`
					Pool struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"pool"`
					Portal struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"portal"`
					Redundancy struct {
						Default     int    `json:"default"`
						Description string `json:"description"`
						Maximum     int    `json:"maximum"`
						Minimum     int    `json:"minimum"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"redundancy"`
					Saferemove struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"saferemove"`
					SaferemoveThroughput struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"saferemove_throughput"`
					Server struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"server"`
					Server2 struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Requires    string `json:"requires"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"server2"`
					Share struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"share"`
					Shared struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"shared"`
					Smbversion struct {
						Description string   `json:"description"`
						Enum        []string `json:"enum"`
						Optional    int      `json:"optional"`
						Type        string   `json:"type"`
					} `json:"smbversion"`
					Sparse struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"sparse"`
					Storage struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"storage"`
					Subdir struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"subdir"`
					TaggedOnly struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"tagged_only"`
					Target struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"target"`
					Thinpool struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"thinpool"`
					Transport struct {
						Description string   `json:"description"`
						Enum        []string `json:"enum"`
						Optional    int      `json:"optional"`
						Type        string   `json:"type"`
					} `json:"transport"`
					Type struct {
						Description string   `json:"description"`
						Enum        []string `json:"enum"`
						Type        string   `json:"type"`
					} `json:"type"`
					Username struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"username"`
					Vgname struct {
						Description string `json:"description"`
						Format      string `json:"format"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"vgname"`
					Volume struct {
						Description string `json:"description"`
						Optional    int    `json:"optional"`
						Type        string `json:"type"`
						Typetext    string `json:"typetext"`
					} `json:"volume"`
				} `json:"properties"`
				Type string `json:"type"`
			} `json:"parameters"`
			Permissions struct {
				Check []interface{} `json:"check"`
			} `json:"permissions"`
			Protected int `json:"protected"`
			Returns   struct {
				Type string `json:"type"`
			} `json:"returns"`
		} `json:"POST"`
	} `json:"info"`
	Leaf int    `json:"leaf"`
	Path string `json:"path"`
	Text string `json:"text"`
}

type Access struct {
	Children []struct {
		Children []struct {
			Children []struct {
				Info struct {
					GET struct {
						Allowtoken  int    `json:"allowtoken"`
						Description string `json:"description"`
						Method      string `json:"method"`
						Name        string `json:"name"`
						Parameters  struct {
							AdditionalProperties int `json:"additionalProperties"`
							Properties           struct {
								Userid struct {
									Description string `json:"description"`
									Format      string `json:"format"`
									MaxLength   int    `json:"maxLength"`
									Type        string `json:"type"`
									Typetext    string `json:"typetext"`
								} `json:"userid"`
							} `json:"properties"`
						} `json:"parameters"`
						Permissions struct {
							Check []interface{} `json:"check"`
						} `json:"permissions"`
						Protected int `json:"protected"`
						Returns   struct {
							AdditionalProperties int `json:"additionalProperties"`
							Properties           struct {
								Realm struct {
									Description string   `json:"description"`
									Enum        []string `json:"enum"`
									Optional    int      `json:"optional"`
									Type        string   `json:"type"`
								} `json:"realm"`
								User struct {
									Description string   `json:"description"`
									Enum        []string `json:"enum"`
									Optional    int      `json:"optional"`
									Type        string   `json:"type"`
								} `json:"user"`
							} `json:"properties"`
							Type string `json:"type"`
						} `json:"returns"`
					} `json:"GET"`
				} `json:"info"`
				Leaf     int    `json:"leaf"`
				Path     string `json:"path"`
				Text     string `json:"text"`
				Children []struct {
					Info struct {
						DELETE struct {
							Allowtoken  int    `json:"allowtoken"`
							Description string `json:"description"`
							Method      string `json:"method"`
							Name        string `json:"name"`
							Parameters  struct {
								AdditionalProperties int `json:"additionalProperties"`
								Properties           struct {
									Tokenid struct {
										Description string `json:"description"`
										Pattern     string `json:"pattern"`
										Type        string `json:"type"`
									} `json:"tokenid"`
									Userid struct {
										Description string `json:"description"`
										Format      string `json:"format"`
										MaxLength   int    `json:"maxLength"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"userid"`
								} `json:"properties"`
							} `json:"parameters"`
							Permissions struct {
								Check []interface{} `json:"check"`
							} `json:"permissions"`
							Protected int `json:"protected"`
							Returns   struct {
								Type string `json:"type"`
							} `json:"returns"`
						} `json:"DELETE"`
						GET struct {
							Allowtoken  int    `json:"allowtoken"`
							Description string `json:"description"`
							Method      string `json:"method"`
							Name        string `json:"name"`
							Parameters  struct {
								AdditionalProperties int `json:"additionalProperties"`
								Properties           struct {
									Tokenid struct {
										Description string `json:"description"`
										Pattern     string `json:"pattern"`
										Type        string `json:"type"`
									} `json:"tokenid"`
									Userid struct {
										Description string `json:"description"`
										Format      string `json:"format"`
										MaxLength   int    `json:"maxLength"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"userid"`
								} `json:"properties"`
							} `json:"parameters"`
							Permissions struct {
								Check []interface{} `json:"check"`
							} `json:"permissions"`
							Returns struct {
								Properties struct {
									Comment struct {
										Optional int    `json:"optional"`
										Type     string `json:"type"`
									} `json:"comment"`
									Expire struct {
										Default     string `json:"default"`
										Description string `json:"description"`
										Minimum     int    `json:"minimum"`
										Optional    int    `json:"optional"`
										Type        string `json:"type"`
									} `json:"expire"`
									Privsep struct {
										Default     int    `json:"default"`
										Description string `json:"description"`
										Optional    int    `json:"optional"`
										Type        string `json:"type"`
									} `json:"privsep"`
								} `json:"properties"`
								Type string `json:"type"`
							} `json:"returns"`
						} `json:"GET"`
						POST struct {
							Allowtoken  int    `json:"allowtoken"`
							Description string `json:"description"`
							Method      string `json:"method"`
							Name        string `json:"name"`
							Parameters  struct {
								AdditionalProperties int `json:"additionalProperties"`
								Properties           struct {
									Comment struct {
										Optional int    `json:"optional"`
										Type     string `json:"type"`
										Typetext string `json:"typetext"`
									} `json:"comment"`
									Expire struct {
										Default     string `json:"default"`
										Description string `json:"description"`
										Minimum     int    `json:"minimum"`
										Optional    int    `json:"optional"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"expire"`
									Privsep struct {
										Default     int    `json:"default"`
										Description string `json:"description"`
										Optional    int    `json:"optional"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"privsep"`
									Tokenid struct {
										Description string `json:"description"`
										Pattern     string `json:"pattern"`
										Type        string `json:"type"`
									} `json:"tokenid"`
									Userid struct {
										Description string `json:"description"`
										Format      string `json:"format"`
										MaxLength   int    `json:"maxLength"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"userid"`
								} `json:"properties"`
							} `json:"parameters"`
							Permissions struct {
								Check []interface{} `json:"check"`
							} `json:"permissions"`
							Protected int `json:"protected"`
							Returns   struct {
								AdditionalProperties int `json:"additionalProperties"`
								Properties           struct {
									Info struct {
										Properties struct {
											Comment struct {
												Optional int    `json:"optional"`
												Type     string `json:"type"`
											} `json:"comment"`
											Expire struct {
												Default     string `json:"default"`
												Description string `json:"description"`
												Minimum     int    `json:"minimum"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
											} `json:"expire"`
											Privsep struct {
												Default     int    `json:"default"`
												Description string `json:"description"`
												Optional    int    `json:"optional"`
												Type        string `json:"type"`
											} `json:"privsep"`
										} `json:"properties"`
										Type string `json:"type"`
									} `json:"info"`
									Value struct {
										Description string `json:"description"`
										Type        string `json:"type"`
									} `json:"value"`
								} `json:"properties"`
								Type string `json:"type"`
							} `json:"returns"`
						} `json:"POST"`
						PUT struct {
							Allowtoken  int    `json:"allowtoken"`
							Description string `json:"description"`
							Method      string `json:"method"`
							Name        string `json:"name"`
							Parameters  struct {
								AdditionalProperties int `json:"additionalProperties"`
								Properties           struct {
									Comment struct {
										Optional int    `json:"optional"`
										Type     string `json:"type"`
										Typetext string `json:"typetext"`
									} `json:"comment"`
									Expire struct {
										Default     string `json:"default"`
										Description string `json:"description"`
										Minimum     int    `json:"minimum"`
										Optional    int    `json:"optional"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"expire"`
									Privsep struct {
										Default     int    `json:"default"`
										Description string `json:"description"`
										Optional    int    `json:"optional"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"privsep"`
									Tokenid struct {
										Description string `json:"description"`
										Pattern     string `json:"pattern"`
										Type        string `json:"type"`
									} `json:"tokenid"`
									Userid struct {
										Description string `json:"description"`
										Format      string `json:"format"`
										MaxLength   int    `json:"maxLength"`
										Type        string `json:"type"`
										Typetext    string `json:"typetext"`
									} `json:"userid"`
								} `json:"properties"`
							} `json:"parameters"`
							Permissions struct {
								Check []interface{} `json:"check"`
							} `json:"permissions"`
							Protected int `json:"protected"`
							Returns   struct {
								Description string `json:"description"`
								Properties  struct {
									Comment struct {
										Optional int    `json:"optional"`
										Type     string `json:"type"`
									} `json:"comment"`
									Expire struct {
										Default     string `json:"default"`
										Description string `json:"description"`
										Minimum     int    `json:"minimum"`
										Optional    int    `json:"optional"`
										Type        string `json:"type"`
									} `json:"expire"`
									Privsep struct {
										Default     int    `json:"default"`
										Description string `json:"description"`
										Optional    int    `json:"optional"`
										Type        string `json:"type"`
									} `json:"privsep"`
								} `json:"properties"`
								Type string `json:"type"`
							} `json:"returns"`
						} `json:"PUT"`
					} `json:"info"`
					Leaf int    `json:"leaf"`
					Path string `json:"path"`
					Text string `json:"text"`
				} `json:"children,omitempty"`
			} `json:"children"`
			Info struct {
				DELETE struct {
					Allowtoken  int    `json:"allowtoken"`
					Description string `json:"description"`
					Method      string `json:"method"`
					Name        string `json:"name"`
					Parameters  struct {
						AdditionalProperties int `json:"additionalProperties"`
						Properties           struct {
							Userid struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								MaxLength   int    `json:"maxLength"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"userid"`
						} `json:"properties"`
					} `json:"parameters"`
					Permissions struct {
						Check []interface{} `json:"check"`
					} `json:"permissions"`
					Protected int `json:"protected"`
					Returns   struct {
						Type string `json:"type"`
					} `json:"returns"`
				} `json:"DELETE"`
				GET struct {
					Allowtoken  int    `json:"allowtoken"`
					Description string `json:"description"`
					Method      string `json:"method"`
					Name        string `json:"name"`
					Parameters  struct {
						AdditionalProperties int `json:"additionalProperties"`
						Properties           struct {
							Userid struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								MaxLength   int    `json:"maxLength"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"userid"`
						} `json:"properties"`
					} `json:"parameters"`
					Permissions struct {
						Check []interface{} `json:"check"`
					} `json:"permissions"`
					Returns struct {
						AdditionalProperties int `json:"additionalProperties"`
						Properties           struct {
							Comment struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"comment"`
							Email struct {
								Format   string `json:"format"`
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"email"`
							Enable struct {
								Default     int    `json:"default"`
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
							} `json:"enable"`
							Expire struct {
								Description string `json:"description"`
								Minimum     int    `json:"minimum"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
							} `json:"expire"`
							Firstname struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"firstname"`
							Groups struct {
								Items struct {
									Format string `json:"format"`
									Type   string `json:"type"`
								} `json:"items"`
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"groups"`
							Keys struct {
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
							} `json:"keys"`
							Lastname struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"lastname"`
							Tokens struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"tokens"`
						} `json:"properties"`
						Type string `json:"type"`
					} `json:"returns"`
				} `json:"GET"`
				PUT struct {
					Allowtoken  int    `json:"allowtoken"`
					Description string `json:"description"`
					Method      string `json:"method"`
					Name        string `json:"name"`
					Parameters  struct {
						AdditionalProperties int `json:"additionalProperties"`
						Properties           struct {
							Append struct {
								Optional int    `json:"optional"`
								Requires string `json:"requires"`
								Type     string `json:"type"`
								Typetext string `json:"typetext"`
							} `json:"append"`
							Comment struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
								Typetext string `json:"typetext"`
							} `json:"comment"`
							Email struct {
								Format   string `json:"format"`
								Optional int    `json:"optional"`
								Type     string `json:"type"`
								Typetext string `json:"typetext"`
							} `json:"email"`
							Enable struct {
								Default     int    `json:"default"`
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"enable"`
							Expire struct {
								Description string `json:"description"`
								Minimum     int    `json:"minimum"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"expire"`
							Firstname struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
								Typetext string `json:"typetext"`
							} `json:"firstname"`
							Groups struct {
								Format   string `json:"format"`
								Optional int    `json:"optional"`
								Type     string `json:"type"`
								Typetext string `json:"typetext"`
							} `json:"groups"`
							Keys struct {
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"keys"`
							Lastname struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
								Typetext string `json:"typetext"`
							} `json:"lastname"`
							Userid struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								MaxLength   int    `json:"maxLength"`
								Type        string `json:"type"`
								Typetext    string `json:"typetext"`
							} `json:"userid"`
						} `json:"properties"`
					} `json:"parameters"`
					Permissions struct {
						Check []interface{} `json:"check"`
					} `json:"permissions"`
					Protected int `json:"protected"`
					Returns   struct {
						Type string `json:"type"`
					} `json:"returns"`
				} `json:"PUT"`
			} `json:"info"`
			Leaf int    `json:"leaf"`
			Path string `json:"path"`
			Text string `json:"text"`
		} `json:"children,omitempty"`
		Info struct {
			GET struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Enabled struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"enabled"`
						Full struct {
							Default     int    `json:"default"`
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"full"`
					} `json:"properties"`
				} `json:"parameters"`
				Permissions struct {
					Description string `json:"description"`
					User        string `json:"user"`
				} `json:"permissions"`
				Returns struct {
					Items struct {
						Properties struct {
							Comment struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"comment"`
							Email struct {
								Format   string `json:"format"`
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"email"`
							Enable struct {
								Default     int    `json:"default"`
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
							} `json:"enable"`
							Expire struct {
								Description string `json:"description"`
								Minimum     int    `json:"minimum"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
							} `json:"expire"`
							Firstname struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"firstname"`
							Groups struct {
								Format   string `json:"format"`
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"groups"`
							Keys struct {
								Description string `json:"description"`
								Optional    int    `json:"optional"`
								Type        string `json:"type"`
							} `json:"keys"`
							Lastname struct {
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"lastname"`
							Tokens struct {
								Items struct {
									Properties struct {
										Comment struct {
											Optional int    `json:"optional"`
											Type     string `json:"type"`
										} `json:"comment"`
										Expire struct {
											Default     string `json:"default"`
											Description string `json:"description"`
											Minimum     int    `json:"minimum"`
											Optional    int    `json:"optional"`
											Type        string `json:"type"`
										} `json:"expire"`
										Privsep struct {
											Default     int    `json:"default"`
											Description string `json:"description"`
											Optional    int    `json:"optional"`
											Type        string `json:"type"`
										} `json:"privsep"`
										Tokenid struct {
											Description string `json:"description"`
											Pattern     string `json:"pattern"`
											Type        string `json:"type"`
										} `json:"tokenid"`
									} `json:"properties"`
									Type string `json:"type"`
								} `json:"items"`
								Optional int    `json:"optional"`
								Type     string `json:"type"`
							} `json:"tokens"`
							Userid struct {
								Description string `json:"description"`
								Format      string `json:"format"`
								MaxLength   int    `json:"maxLength"`
								Type        string `json:"type"`
							} `json:"userid"`
						} `json:"properties"`
						Type string `json:"type"`
					} `json:"items"`
					Links []struct {
						Href string `json:"href"`
						Rel  string `json:"rel"`
					} `json:"links"`
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"GET"`
			POST struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Comment struct {
							Optional int    `json:"optional"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"comment"`
						Email struct {
							Format   string `json:"format"`
							Optional int    `json:"optional"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"email"`
						Enable struct {
							Default     int    `json:"default"`
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"enable"`
						Expire struct {
							Description string `json:"description"`
							Minimum     int    `json:"minimum"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"expire"`
						Firstname struct {
							Optional int    `json:"optional"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"firstname"`
						Groups struct {
							Format   string `json:"format"`
							Optional int    `json:"optional"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"groups"`
						Keys struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"keys"`
						Lastname struct {
							Optional int    `json:"optional"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"lastname"`
						Password struct {
							Description string `json:"description"`
							MaxLength   int    `json:"maxLength"`
							MinLength   int    `json:"minLength"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"password"`
						Userid struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							MaxLength   int    `json:"maxLength"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"userid"`
					} `json:"properties"`
				} `json:"parameters"`
				Permissions struct {
					Check       []interface{} `json:"check"`
					Description string        `json:"description"`
				} `json:"permissions"`
				Protected int `json:"protected"`
				Returns   struct {
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"POST"`
		} `json:"info,omitempty"`
		Leaf int    `json:"leaf"`
		Path string `json:"path"`
		Text string `json:"text"`
	} `json:"children"`
	Info struct {
		GET struct {
			Allowtoken  int    `json:"allowtoken"`
			Description string `json:"description"`
			Method      string `json:"method"`
			Name        string `json:"name"`
			Parameters  struct {
				AdditionalProperties int `json:"additionalProperties"`
			} `json:"parameters"`
			Permissions struct {
				User string `json:"user"`
			} `json:"permissions"`
			Returns struct {
				Items struct {
					Properties struct {
						Subdir struct {
							Type string `json:"type"`
						} `json:"subdir"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"items"`
				Links []struct {
					Href string `json:"href"`
					Rel  string `json:"rel"`
				} `json:"links"`
				Type string `json:"type"`
			} `json:"returns"`
		} `json:"GET"`
	} `json:"info"`
	Leaf int    `json:"leaf"`
	Path string `json:"path"`
	Text string `json:"text"`
}

type Pools struct {
	Children []struct {
		Info struct {
			DELETE struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Poolid struct {
							Format   string `json:"format"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"poolid"`
					} `json:"properties"`
				} `json:"parameters"`
				Permissions struct {
					Check       []interface{} `json:"check"`
					Description string        `json:"description"`
				} `json:"permissions"`
				Protected int `json:"protected"`
				Returns   struct {
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"DELETE"`
			GET struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Poolid struct {
							Format   string `json:"format"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"poolid"`
					} `json:"properties"`
				} `json:"parameters"`
				Permissions struct {
					Check []interface{} `json:"check"`
				} `json:"permissions"`
				Returns struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Comment struct {
							Optional int    `json:"optional"`
							Type     string `json:"type"`
						} `json:"comment"`
						Members struct {
							Items struct {
								AdditionalProperties int `json:"additionalProperties"`
								Properties           struct {
									ID struct {
										Type string `json:"type"`
									} `json:"id"`
									Node struct {
										Type string `json:"type"`
									} `json:"node"`
									Storage struct {
										Optional int    `json:"optional"`
										Type     string `json:"type"`
									} `json:"storage"`
									Type struct {
										Enum []string `json:"enum"`
										Type string   `json:"type"`
									} `json:"type"`
									Vmid struct {
										Optional int    `json:"optional"`
										Type     string `json:"type"`
									} `json:"vmid"`
								} `json:"properties"`
								Type string `json:"type"`
							} `json:"items"`
							Type string `json:"type"`
						} `json:"members"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"GET"`
			PUT struct {
				Allowtoken  int    `json:"allowtoken"`
				Description string `json:"description"`
				Method      string `json:"method"`
				Name        string `json:"name"`
				Parameters  struct {
					AdditionalProperties int `json:"additionalProperties"`
					Properties           struct {
						Comment struct {
							Optional int    `json:"optional"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"comment"`
						Delete struct {
							Description string `json:"description"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"delete"`
						Poolid struct {
							Format   string `json:"format"`
							Type     string `json:"type"`
							Typetext string `json:"typetext"`
						} `json:"poolid"`
						Storage struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"storage"`
						Vms struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Optional    int    `json:"optional"`
							Type        string `json:"type"`
							Typetext    string `json:"typetext"`
						} `json:"vms"`
					} `json:"properties"`
				} `json:"parameters"`
				Permissions struct {
					Check       []interface{} `json:"check"`
					Description string        `json:"description"`
				} `json:"permissions"`
				Protected int `json:"protected"`
				Returns   struct {
					Type string `json:"type"`
				} `json:"returns"`
			} `json:"PUT"`
		} `json:"info"`
		Leaf int    `json:"leaf"`
		Path string `json:"path"`
		Text string `json:"text"`
	} `json:"children"`
	Info struct {
		GET struct {
			Allowtoken  int    `json:"allowtoken"`
			Description string `json:"description"`
			Method      string `json:"method"`
			Name        string `json:"name"`
			Parameters  struct {
				AdditionalProperties int `json:"additionalProperties"`
			} `json:"parameters"`
			Permissions struct {
				Description string `json:"description"`
				User        string `json:"user"`
			} `json:"permissions"`
			Returns struct {
				Items struct {
					Properties struct {
						Poolid struct {
							Type string `json:"type"`
						} `json:"poolid"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"items"`
				Links []struct {
					Href string `json:"href"`
					Rel  string `json:"rel"`
				} `json:"links"`
				Type string `json:"type"`
			} `json:"returns"`
		} `json:"GET"`
		POST struct {
			Allowtoken  int    `json:"allowtoken"`
			Description string `json:"description"`
			Method      string `json:"method"`
			Name        string `json:"name"`
			Parameters  struct {
				AdditionalProperties int `json:"additionalProperties"`
				Properties           struct {
					Comment struct {
						Optional int    `json:"optional"`
						Type     string `json:"type"`
						Typetext string `json:"typetext"`
					} `json:"comment"`
					Poolid struct {
						Format   string `json:"format"`
						Type     string `json:"type"`
						Typetext string `json:"typetext"`
					} `json:"poolid"`
				} `json:"properties"`
			} `json:"parameters"`
			Permissions struct {
				Check []interface{} `json:"check"`
			} `json:"permissions"`
			Protected int `json:"protected"`
			Returns   struct {
				Type string `json:"type"`
			} `json:"returns"`
		} `json:"POST"`
	} `json:"info"`
	Leaf int    `json:"leaf"`
	Path string `json:"path"`
	Text string `json:"text"`
}

type Version struct {
	Info struct {
		GET struct {
			Allowtoken  int    `json:"allowtoken"`
			Description string `json:"description"`
			Method      string `json:"method"`
			Name        string `json:"name"`
			Parameters  struct {
				AdditionalProperties int `json:"additionalProperties"`
			} `json:"parameters"`
			Permissions struct {
				User string `json:"user"`
			} `json:"permissions"`
			Returns struct {
				Properties struct {
					Release struct {
						Type string `json:"type"`
					} `json:"release"`
					Repoid struct {
						Type string `json:"type"`
					} `json:"repoid"`
					Version struct {
						Type string `json:"type"`
					} `json:"version"`
				} `json:"properties"`
				Type string `json:"type"`
			} `json:"returns"`
		} `json:"GET"`
	} `json:"info"`
	Leaf int    `json:"leaf"`
	Path string `json:"path"`
	Text string `json:"text"`
}
